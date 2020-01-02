package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	op := flag.Int("option", 256, "sha option(256, 384, 512")
	in := flag.String("input", "0", "input string")

	flag.Parse()

	switch *op {
	case 256:
		fmt.Println(sha256.Sum256([]byte(*in)))
	case 384:
		fmt.Println(sha512.Sum384([]byte(*in)))
	case 512:
		fmt.Println(sha512.Sum512([]byte(*in)))
	default:
		fmt.Println("????")
	}
}
