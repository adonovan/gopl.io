// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// ex 1.2 solution
// modify echo program to print the index and value of each of it's
// arguments, one per line.

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, x := range os.Args[1:] {
		fmt.Println(i, x)
	}
}
