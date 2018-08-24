// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// ex1.3
// benchmark the difference between two of the echo implementations from
// this section.

/*

$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/guidorice/gopl.io/ch1/ex1_3
BenchmarkEcho1-4   	10000000	       118 ns/op
BenchmarkEcho3-4   	20000000	        68.9 ns/op
PASS
ok  	github.com/guidorice/gopl.io/ch1/ex1_3	2.774s

 */
package ex1_3

import (
	"strings"
)

// echo1/main uses string concatenation (expect to be slow)
func echo1(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

// echo3/main uses strings.Join and a slice (expect to be faster)
func echo3(args []string) string {
	return strings.Join(args[1:], " ")
}
