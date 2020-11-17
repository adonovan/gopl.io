// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

//!+
func main() {
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("Index: ", i, "Value: ", os.Args[i], time.Since(start).Seconds())
	}

	//start := time.Now()
	//fmt.Println(strings.Join(os.Args[1:], " "), time.Since(start).Seconds())
}

//!-
