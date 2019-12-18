// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

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
		fmt.Println("Enter a filename.")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			if isDuplicated(f, counts) {
				fmt.Println(arg)
			}
			f.Close()
		}
	}

}

func isDuplicated(f *os.File, counts map[string]int) bool {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()]++; counts[input.Text()] > 1 {
			return true
		}
	}

	return false
}
