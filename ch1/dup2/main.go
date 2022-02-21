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
	dupLineFileName := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupLineFileName)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dupLineFileName)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
		//delete(counts, line)
		//fmt.Println("deleted ",counts)
	}
	// for k := range counts {
	// 	delete(counts, k)
	// }
	
	for dlfn, dlf := range dupLineFileName {
		fmt.Println(dlfn, dlf)
	}
}

func countLines(f *os.File, counts map[string]int, dupLineFileName map[string]map[string]int) {
	input := bufio.NewScanner(f)
	//fmt.Println(f.Name())
	for input.Scan() {
		if(input.Text() == "end") { break }
		counts[input.Text()]++
		// counts[input.Text()] = 
	}
	for _, n := range counts {
		if n > 1 {
			fmt.Println("in if ",n)
		dupLineFileName[f.Name()] = counts
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
