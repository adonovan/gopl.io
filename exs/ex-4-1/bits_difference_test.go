// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

var c1, c2 [32]byte

func init() {
	c1 = sha256.Sum256([]byte("unodostresquatro"))
	c2 = sha256.Sum256([]byte("UNODOSTRESQUATRO"))
}

// -- Alternative implementations --

func AlternativeBitsDifference(h1, h2 *[sha256.Size]byte) int {
	n := 0
	for i := range h1 {
		for b := h1[i] ^ h2[i]; b != 0; b &= b - 1 {
			n++
		}
	}
	return n
}

func assertEquals(x, y int) {
	if x != y {
		panic(fmt.Sprintf("%d != %d", x, y))
	}
}

func TestSplit(t *testing.T) {
	assertEquals(AlternativeBitsDifference(&c1, &c2), 139)
	assertEquals(BitsDifference(&c1, &c2), 139)
}

// -- Benchmarks --

func BenchmarkAlternativeBitsDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlternativeBitsDifference(&c1, &c2)
	}
}

func BenchmarkBitsDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitsDifference(&c1, &c2)
	}
}

// Go 1.6, 2.67GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkBitsDifference-4                  200000000         6.30 ns/op
// BenchmarkAlternativeBitsDifference-4                  300000000         4.15 ns/op
// BenchmarkPopCountByClearing-4        30000000         45.2 ns/op
// BenchmarkPopCountByShifting-4        10000000        153 ns/op
//
// Go 1.6, 2.5GHz Intel Core i5
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkBitsDifference-4                  200000000         7.52 ns/op
// BenchmarkAlternativeBitsDifference-4                  500000000         3.36 ns/op
// BenchmarkPopCountByClearing-4        50000000         34.3 ns/op
// BenchmarkPopCountByShifting-4        20000000        108 ns/op
//
// Go 1.7, 3.5GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkBitsDifference-12                 2000000000        0.28 ns/op
// BenchmarkAlternativeBitsDifference-12                 2000000000        0.27 ns/op
// BenchmarkPopCountByClearing-12       100000000        18.5 ns/op
// BenchmarkPopCountByShifting-12       20000000         70.1 ns/op
