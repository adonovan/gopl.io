package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	dupfiles := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, dupfiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				// %v : the value in a default format
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dupfiles)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("%s\t%v\n", line, dupfiles[line])
		}
	}

}

func countLines(f *os.File, counts map[string]int, dupfiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		dupfiles[input.Text()] = append(dupfiles[input.Text()], f.Name())
	}
}