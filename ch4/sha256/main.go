// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
//!+
package main

import (
	"crypto/sha256"
	"flag"
	"fmt"

	"gopl.io/ch4/sha256/sha512"
)
var algo = flag.String("a", "SHA256", "algorithm name")
var text = flag.String("t", "Hello World!", "Text which needs to be hashed")
var result string
//var algoByte [32]byte = [32]byte{}
var textByte []byte = []byte{}
func main() {
	flag.Parse()
	textByte = []byte(*text)

	if(*algo != "" || *text != ""){
		result = sha512.CommandLineHash(*algo, textByte)
	}else{
		byteArray := sha256.Sum256([]byte(textByte))
		result = string(byteArray[:])
	}

	fmt.Printf("%x \n", result)
}

//!-
