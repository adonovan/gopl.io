package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)

		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	} else {
		for _, arg := range files {
			fileLineCounts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			defer f.Close()

			if countLines(f, fileLineCounts) {
				fmt.Println("dup file name:", arg)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) bool {
	input := bufio.NewScanner(f)
	text := input.Text()
	isDup := false
	for input.Scan() {
		counts[text]++
		if counts[text] > 1 {
			isDup = true
		}
	}
	return isDup
}
