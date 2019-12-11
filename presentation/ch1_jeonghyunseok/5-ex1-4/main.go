// 1.4 dup2를 수정해 중복된 줄이 있는 파일명을 모두 출력하라.
// start with dup2
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
		countLines(os.Args[0]+":", os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(arg, f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(fileName string, f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[fileName+":"+input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-

/*
go build -o ex1-4.exe
ex1-4.exe a.txt b.txt c.txt

*/

// 2       a.txt:ago
// 2       b.txt:beyond
// 3       c.txt:cup
// 2       c.txt:cloud
// 3       a.txt:about
