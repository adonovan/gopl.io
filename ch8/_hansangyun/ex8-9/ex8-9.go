package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

/*
연습문제 8.9

각 root 디렉토리별 총합을 계산하고 주기적으로 출력하는 du 버전을 작성하라.

go run ex8-9/ex8-9.go -v ~/Dropbox ~/Downloads
 */

var verbose = flag.Bool("v", false, "show verbose progress messages")

type SizeResponse struct {
	root int
	size int64
}


func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	s := make(chan SizeResponse)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root, &n, i, s)
	}
	go func() {
		n.Wait()
		close(s)

	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(10 * time.Millisecond)
	}

	// 입력한 root path의 갯수만큼 생성, 루트별로 따로 카운팅
	/*
	root1 ==> counting...
	root2 ==> counting...
	...
	*/
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))

loop:
	for {
		select {
		case sizeResponse, ok := <-s:
			if !ok {
				break loop
			}
			nfiles[sizeResponse.root]++
			nbytes[sizeResponse.root] += sizeResponse.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}
	printDiskUsage(roots, nfiles, nbytes)
}


func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	fmt.Println("================================")
	for i, r := range roots {
		fmt.Printf("%10d files  %.3f GB under %s\n", nfiles[i], float64(nbytes[i])/1e9, r)
	}
	fmt.Println("================================\n\n")
}

func walkDir(dir string, n *sync.WaitGroup, root int, s chan<- SizeResponse) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, root, s)
		} else {
			s <- SizeResponse{root, entry.Size()}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() {<-sema}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "01-du1: %v\n", err)
		return nil
	}
	return entries
}