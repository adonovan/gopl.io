package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

const (
	// Size is the size, in bytes, of a SHA-512 checksum.
	Size = 64

	// Size224 is the size, in bytes, of a SHA-512/224 checksum.
	Size224 = 28

	// Size256 is the size, in bytes, of a SHA-512/256 checksum.
	Size256 = 32

	// Size384 is the size, in bytes, of a SHA-384 checksum.
	Size384 = 48

	// BlockSize is the block size, in bytes, of the SHA-512/224,
	// SHA-512/256, SHA-384 and SHA-512 hash functions.
	BlockSize = 128
)

func cryptoToStringProcess() {
	var shaName string
	flag.StringVar(&shaName, "sha-name", "sha265", "a string sha name")
	flag.Parse()

	if shaName == "sha384" {
		sha := sha512.Sum384([]byte("kimhyegon"))
		fmt.Printf("%x \n", sha)
	} else if shaName == "sha512" {
		sha := sha512.Sum512([]byte("kimhyegon"))
		fmt.Printf("%x \n", sha)

	} else {
		sha := sha256.Sum256([]byte("kimhyegon"))
		fmt.Printf("%x \n", sha)
	}

}

func cryptoToIntProcess() {
	var shaSize int
	flag.IntVar(&shaSize, "sha-size", Size256, "help crypto encoding size")
	flag.Parse()

	if shaSize == Size384 {
		sha := sha512.Sum384([]byte("kimhyegon"))
		fmt.Printf("%x \n", sha)
	} else if shaSize == BlockSize {
		sha := sha512.Sum512([]byte("kimhyegon"))
		fmt.Printf("%x \n", sha)

	} else {
		sha := sha256.Sum256([]byte("kimhyegon"))
		fmt.Printf("%x \n", sha)
	}

}

func main() {

	argument := os.Args
	if len(argument) != 3 {
		fmt.Printf("............. Arguemtn 입력 오류 사용법 : %s -sha-name sha 384 or %s -sha-size 384\n", argument[0], argument[0])
		os.Exit(1)
	}

	if argument[1] == "-sha-name" {
		cryptoToStringProcess()
	} else {
		cryptoToIntProcess()
	}
}
