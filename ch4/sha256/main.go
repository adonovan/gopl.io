// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"fmt"

	"github.com/popcount"
)

//!+
import "crypto/sha256"

func main() {
	c1 := sha256.Sum256([]byte("now is the time"))
	c2 := sha256.Sum256([]byte("Bobs your uncle"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	c3 := [32]uint8{}

	for i := 0; i < len(c1); i++ {
		c3[i] = c1[i] ^ c2[i]
	}

	fmt.Printf("%x\n", c3)

	var count int

	for i := 0; i < len(c3); i++ {
		count = count + popcount.PopCountByte(c3[i])
		fmt.Printf("Popcount of byte %d is (%b) is %d\n", i, c3[i], popcount.PopCountByte(c3[i]))
	}

	fmt.Printf("Population count is %d\n", count)

	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-
