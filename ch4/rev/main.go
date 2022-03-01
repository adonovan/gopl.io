// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"gopl.io/ch4/rev/reverse"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse.Reverse(a[:])
	fmt.Println("reverse string of a : ",a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	t := make([]int, len(s))
	copy(t, s)
	// Rotate s left by two positions.
	reverse.Reverse(s[:2])
	fmt.Println(":2 ",s) // 1,0,2,3,4,5
	reverse.Reverse(s[2:]) 
	fmt.Println(":2 ",s) // 1,0,5,4,3,2
	reverse.Reverse(s) //
	fmt.Println("move left by 2 positions",s) // "[2 3 4 5 0 1]"
	//!-slice
	//var x []*int = &s
	var s1 = [...]int{10, 11, 12, 13, 14, 15}
	fmt.Println("\n", reflect.ValueOf(s1).Kind())
	reverse.ReverseUsingPointer(&s1)
	fmt.Println("ReverseUsingPointer : ", s1)
	reverse.Rotate(t, 2, true)
	fmt.Println("t : ", t)
	reverse.RemoveDuplicatesFromStringSlice([]string{"abc", "abc","ccc", "ddd","ccc","ccc"})
	reverse.ReverseTheCharOfByteSlice([]byte("Hello"))
	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
mainloop:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue mainloop
			}
			ints = append(ints, int(x))
		}
		reverse.Reverse(ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

