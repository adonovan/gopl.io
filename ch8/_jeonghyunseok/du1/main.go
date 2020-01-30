// 디렉토리의 파일들의 사용을 계산해주는 프로그램이다.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// 파일트리를 돌아다녀 보낟.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nf, nb int64) {
	fmt.Printf("%d files, %.1f GB\n", nf, float64(nb)/1e9)
}

// 재귀적으로 파일트리를 돌아다닌다
func walkDir(dir string, fsizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fsizes)
		} else {
			fsizes <- entry.Size()
		}
	}
}

// dirents 는 디렉토리의 entries 들을 리턴해준다.
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

/*
1) dirents 부터 보자
ioutil.ReadDir() 해주면 []os.FileInfo 를 준다. 

	type FileInfo interface {
		Name() string       // base name of the file
		Size() int64        // length in bytes for regular files; system-dependent for others
		Mode() FileMode     // file mode bits
		ModTime() time.Time // modification time
		IsDir() bool        // abbreviation for Mode().IsDir()
		Sys() interface{}   // underlying data source (can return nil)
	}

2) waklDir 을 보자. 
- 일단 dirents() 를 통해서 FileInfo 들을 받아서 루프를 돌린다.
- 디렉토리라면 다시 workdir 을 재귀적으로 돌리고
- 아니면 fsizes 채널로 사이즈 값을 넣어준다. 

3) 이제 main() 을 보자 
flag.Parse() 해주고 flag.Args() 를 해주면 플래그 아닌 파라미터들을 리턴해준다. 

	// Args returns the non-flag command-line arguments.
	func Args() []string { return CommandLine.args }

	아무것도 없다면 디폴트로 그냥 현재폴더 "."
	있다면 goroutine 을 통해서 roots 를 돌면서 workDir 을 해준다. 이게 전부이다. 

	그리고 아래쪽에서 fileSizes 채널로 들어오는 값을 수집

*/