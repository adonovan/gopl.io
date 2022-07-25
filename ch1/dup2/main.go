package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	//founded_files := make([]int, 0)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			res := countLines(f, counts)
			if res == true {
				fmt.Printf("%v\n", arg)
			}

			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) bool {
	input := bufio.NewScanner(f)
	found := false
	//text := ""
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if counts[text] > 2 {
			found = true
		}
	}
	return found
}
