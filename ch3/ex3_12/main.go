/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package main

import (
	"fmt"
	"os"

	"github.com/guidorice/gopl.io/ch3/ex3_12/anagram"
)

func usage() {
	fmt.Printf("Usage: %s word1 word2\n", os.Args[0])
}

func main() {
	if len(os.Args) != 3 {
		usage()
		os.Exit(1)
	}
	a := os.Args[1]
	b := os.Args[2]
	are := anagram.Anagram(a, b)
	fmt.Printf("%s and %s", a, b)
	if are {
		fmt.Printf(" are anagrams!\n")
	} else {
		fmt.Printf(" are not anagrams.\n")
	}
}
