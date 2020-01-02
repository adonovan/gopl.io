package main

import (
	"crypto/sha256"
	"fmt"
)

func diffSHACount(c1, c2 [32]byte) int {
	count := 0

	for i := range c1 {
		isDiff := false
		for j := uint(0); j < 8; j++ {
			bit1 := (c1[i] >> j) & 1
			bit2 := (c2[i] >> j) & 1
			if bit1 != bit2 {
				count++
				isDiff = true
			}
		}
		if isDiff {
			fmt.Printf("%08b\n%08b\n\n", c1[i], c2[i])
		}
	}

	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Println("diff sha count : ", diffSHACount(c1, c2))
}
