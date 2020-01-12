package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	c1Len := len(c1)
	c2Len := len(c2)
	if c1Len != c2Len {
		fmt.Println("둘의 길이가 다름")
		os.Exit(1)
	}

	for i := 0; i < c1Len; i++ {
		if c1[i] != c2[i] {
			fmt.Printf("0x%x : 0x%x \n", c1[i], c2[i])
		}
	}
	fmt.Println()
}
