// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

/* ex 3.10
 * write a non-recursive version of comma, using bytes.Buffer instead of string
 * concatenation.
 */

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"

	"github.com/guidorice/gopl.io/ch3/ex3_10/comma"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		s := os.Args[i]
		fmt.Printf("  %s -> %s \n", s, comma.Comma(s))
	}
}
