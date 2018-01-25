// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// ex1.3
// this would be better implemented using go test -bench
// and BenchmarkXXX(b *testing.B), however the main package is getting in the way

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo3() string {
	return strings.Join(os.Args[1:], " ")
}

func main() {
	start := time.Now()
	n := 100000

	// echo1
	for i := 0; i < n; i++ {
		echo1()
	}
	echo1Elapsed := time.Since(start).Seconds()

	// echo3
	start = time.Now()
	for i := 0; i < n; i++ {
		echo3()
	}
	echo3Elapsed := time.Since(start).Seconds()
	fmt.Println("echo1:", echo1Elapsed)
	fmt.Println("echo3:", echo3Elapsed)
}
