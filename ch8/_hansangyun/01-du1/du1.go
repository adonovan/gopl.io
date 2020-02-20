package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// go run du1 ~/Downloads ~/Dropbox
// 디렉토리의 디스크 사용량을 재귀적으로 순회하면서 사용량을 더한다
func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 파일 사이즈를 계속 읽어오기 위한 채널
	fileSizes := make(chan int64)

	// 메인 고루틴에선 채널의 읽으면서 전체 사이즈를 누적하고
	// 이 고루틴에선 디렉토리를 순회하면서 파일 사이즈를 읽기 위해 재귀적으로 walkDir을 호출한다.
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes) //TODO 여러 디렉토리를 병렬적으로 수행 가능
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

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		// 디렉토리이면 더 깊이~ 아니면 사이즈를 채널에 입력
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "01-du1: %v\n", err)
		return nil
	}
	return entries
}