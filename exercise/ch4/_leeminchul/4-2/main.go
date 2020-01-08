package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
	"unsafe"
)

// bool2int returns 0 if x is false or 1 if x is true.
func bool2int(x bool) int {
	// Avoid branches. In the SSA compiler, this compiles to
	// exactly what you would want it to.
	return int(*(*uint8)(unsafe.Pointer(&x)))
}

func main() {
	var sha256Flag, sha384Flag, sha512Flag bool
	flag.BoolVar(&sha256Flag, "sha256", false, "SHA256 checksum of the input")
	flag.BoolVar(&sha384Flag, "sha384", false, "SHA384 checksum of the input")
	flag.BoolVar(&sha512Flag, "sha512", false, "SHA512 checksum of the input")

	flag.Parse()
	optionCnt := bool2int(sha256Flag) + bool2int(sha384Flag) + bool2int(sha512Flag)

	if 1 < optionCnt {
		_, _ = fmt.Fprintf(os.Stderr, "Invalid parameter: duplicated options\n")
		os.Exit(1)
	}

	if optionCnt == 0 {
		sha256Flag = true
	}

	in := strings.Join(flag.Args(), " ")
	fmt.Printf("%#v\n", in)
	if sha256Flag {
		fmt.Printf("%064x", sha256.Sum256([]byte(in)))
	}
	if sha384Flag {
		fmt.Printf("%064x", sha512.Sum384([]byte(in)))
	}
	if sha512Flag {
		fmt.Printf("%064x", sha512.Sum512([]byte(in)))
	}
}

// go run . x
// go run . -sha256 x
// go run . -sha384 x
// go run . -sha512 x
