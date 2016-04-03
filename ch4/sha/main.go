package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var t = flag.Bool("384", false, "SHA384")
var f = flag.Bool("512", false, "SHA512")

func main() {

	flag.Parse()

	if *t {
		fmt.Printf("SHA384 = %x\n", sha512.Sum384([]byte(os.Args[1])))
	} else if *f {
		fmt.Printf("SHA512 = %x\n", sha512.Sum512([]byte(os.Args[1])))
	} else {
		fmt.Printf("SHA256 = %x\n", sha256.Sum256([]byte(os.Args[1])))
	}
}
