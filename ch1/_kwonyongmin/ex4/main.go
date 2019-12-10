package main

import (
	"bufio"
	"fmt"
	"os"
)

const noFile string = "?#@#$*#$*($&@*(#!)*#&@*(#&!@)!&#(@$^!*(#$$%@"

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, struct {
			name  string
			files map[string][]string
		}{noFile, nil})
	} else {
		contents := make(map[string][]string)
		for _, arg := range files {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, struct {
				name  string
				files map[string][]string
			}{arg, contents})
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, info struct {
	name  string
	files map[string][]string
}) {
	scanner := bufio.NewScanner(f)
	subCounts := make(map[string]int)

	for scanner.Scan() {
		if info.name != noFile {
			if _, ok := info.files[scanner.Text()]; ok {
				if _, ok := subCounts[scanner.Text()]; !ok {
					info.files[scanner.Text()] = append(info.files[scanner.Text()], info.name)
				}
			} else {
				info.files[scanner.Text()] = []string{info.name}
			}
		}
		subCounts[scanner.Text()]++
		counts[scanner.Text()]++
	}
}
