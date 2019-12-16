package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var filename [10]string
	counts := make(map[string]int)
	files := os.Args[1:]
	index := 0
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			filename[index] = files[index]
			countLines(f, counts)
			f.Close()
		}

	}

	index = 0
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s:%d\t%s\n", filename[index], n, line)
		}

		index++
	}
}

//func countLines(f *os.File, counts map[string]int) {
//	input := bufio.NewScanner(f)
//	for input.Scan() {
//		counts[input.Text()]++
//	}
//
//	// NOTE: input.Err() 에서의 잠재적 오류는 무시한다.ㅊㅇ
//}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

	// NOTE: input.Err() 에서의 잠재적 오류는 무시한다.ㅊㅇ
}
