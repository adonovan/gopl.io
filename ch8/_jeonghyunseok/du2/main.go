// du2 uses select and time.Ticker
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// -v 플래그를 사용한다.
var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {

	// 플래그가 있건 없건 flag.Args() 로 플래그 아닌 녀석들 다 가져올 수 있나 보다.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// tick 은 time.Time 을 읽는 채널이고
	// verbose 플래그가 설정되었다면 500ms Tick 을 넣어준다.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(50 * time.Millisecond) // 이렇게 해주면 500ms 마다 tick 에 값을 써준다고 보면 된다.
	}
	var nfiles, nbytes int64

	// 1. close 체크를 해준다.
	// 2. tick 이 들어올때마다 중간 출력을 해준다.
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case t := <-tick:
			fmt.Println("time: ", t)
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

//!-

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

/*
go build -o du2.exe
du2.exe -v D:\study

*/
