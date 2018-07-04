/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// ex 4.1 write a function that counts the number of bits that are different
// in two SHA256 hashes. (see PopCount from section 2.6.2)

// n.b. the bitcount cannot be calculated directly on an uint because the
// Sum256() returns [32]uint8. instead check each array position for equality.

package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

type hash [sha256.Size]uint8

// counter returns the number of ints that are equal in two hash arrays.
func counter(a hash, b hash) int {
	count := 0
	for i, n := range a {
		if b[i] != n {
			count++
		}
	}
	return count
}

func usage() {
	fmt.Printf("Usage: %s string1 string2\n", os.Args[0])
}

func main() {
	if len(os.Args) != 3 {
		usage()
		os.Exit(1)
	}
	a := os.Args[1]
	b := os.Args[2]
	ah := sha256.Sum256([]byte(a))
	bh := sha256.Sum256([]byte(b))
	n := counter(hash(ah), hash(bh))
	fmt.Printf("sha256 hashes differ by %d bits\n", n)
}
