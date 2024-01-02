// Copyright © 2024 Victor N. Skurikhin.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
// Напишите функцию, которая подсчитывает количество битов, которые отличаются в двух хэшах SHA256.

package main

import (
	"crypto/sha256" //!+
	"fmt"
)

var pc [8]byte

func init() {
	for i := range pc {
		pc[i] = 1 << i
	}
}

func compareByIndex(x, y byte, i int) byte {
	return ((pc[i] & x) ^ (pc[i] & y)) >> i
}

func bitsDifference(x, y byte) int {
	var sum = 0
	for i := 0; i < 8; i++ {
		sum += int(compareByIndex(x, y, i))
	}
	return sum
}

func BitsDifference(x, y *[32]byte) int {
	var sum = 0
	for i, z := range x {
		sum += bitsDifference(z, y[i])
	}
	return sum
}

func main() {
	c1 := sha256.Sum256([]byte("unodostresquatro"))
	c2 := sha256.Sum256([]byte("UNODOSTRESQUATRO"))
	fmt.Printf("count = %d\n", BitsDifference(&c1, &c2))
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
