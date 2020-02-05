// du3 은 병렬적으로 계산해준다
// semaphore used

package main

import (
	"flag"
	"sync"
	"time"
	"os"
	"fmt"
	"path/filepath"
	"io/ioutil"
	
)

var vFlag = flag.Bool("v", false, "show verbose progress message")

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}

	}

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
	loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}	
			printDiskUsage(nfiles, nbytes)

}

func printDiskUsage(nfiles, nbytes int64){
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)

		} else {
			fileSizes <-entry.Size()
		}
	}
}

var sema = make (chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

/*
go build -o du3.exe
du3 --help
du3 -v c:

*/