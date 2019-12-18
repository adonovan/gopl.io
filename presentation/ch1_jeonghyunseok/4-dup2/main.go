// Dup2 check lines more than once
// read from stdin or files
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sameLineCounts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, sameLineCounts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, sameLineCounts)
		}
	}

	for lineText, n := range sameLineCounts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, lineText)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// ignore input.Err()
}

/*
go build -o dup2.exe
dup2 tmp.txt

*/
// 2       bye
// 3       good
// 2       hello
