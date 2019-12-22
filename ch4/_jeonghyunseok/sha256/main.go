// The sha256 command computes the SHA256 hash (an array) of a string
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println("--")

	c3 := sha256.Sum224([]byte("x"))
	c4 := sha256.Sum224([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c3, c4, c3 == c4, c3)

}

// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
// false
// [32]uint8
// --
// 54a2f7f92a5f975d8096af77a126edda7da60c5aa872ef1b871701ae
// f00bdeb2cd9da240a57c951fdf1bcba509fd0cd83c5e5ad9a669de12
// false
// [28]uint8
