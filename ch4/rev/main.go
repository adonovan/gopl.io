// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])                         // The reverse function requires a slice
	fmt.Printf("Rev with Slice: %d\n", a) // "[5 4 3 2 1 0]"

	//Excersize 4.3
	reverse2(&a)                            // the reverse2 function requires a pointer to the whole array
	fmt.Printf("ReRev with Array: %d\n", a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	fmt.Printf("Rot 1\n")
	reverse(s[:2])
	fmt.Printf("Rot 2\n")
	reverse(s[2:])
	fmt.Printf("Rot 3\n")
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	// Excersize 4.4, I think I cheated by simple calling the reverse function
	// 3 times from the rotate function
	rotate(s)

	fmt.Println(s) // "[4 5 0 1 2 3]"
	//!-slice

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("i=%d, j=%d\n", i, j)
		s[i], s[j] = s[j], s[i]
		fmt.Println(s)
	}
	fmt.Printf("\n")
}

// Doug's reverse function with an array instead of a slice

func reverse2(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("i=%d, j=%d\n", i, j)

		s[i], s[j] = s[j], s[i]
	}
	fmt.Printf("\n")
}

//Doug's function to rotate left 2 positions
func rotate(s []int) {
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)

}

//!-rev
