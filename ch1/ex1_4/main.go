// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

// ex1.4 modify dup2 to print the names of all the files in which each
// duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	// map of duplicate line to set of files it occurredIn
	occurredIn := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, occurredIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, occurredIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			var files []string
			for key := range occurredIn[line] {
				files = append(files, key)
			}
			fmt.Print("[", strings.Join(files, ", "), "]")
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int,
	occurredIn map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if occurredIn[line] == nil {
			occurredIn[line] = make(map[string]bool)
		}
		occurredIn[line][f.Name()] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
