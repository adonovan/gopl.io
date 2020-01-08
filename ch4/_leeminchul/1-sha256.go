package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n", c1)
	fmt.Printf("%x\n", c2)
	fmt.Printf("%t\n", c1 == c2) // %t: the word true or false
	fmt.Printf("%T\n", c1)       // %T: type of the value
}

/* output:
2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
false
[32]uint8
*/
