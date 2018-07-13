/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

 /*
 $ go run main.go -help
  -algorithm string
        hash algorithm: sha256, sha384 or sha512 (default "sha256")
  -help
        display help

$ go run main.go alex
4135aa9dc1b842a653dea846903ddb95bfb8c5a10c504a7fa16e10bc31d1fdf0

$ go run main.go -algorithm=sha384 alex
8d47762a83550b50e28f66fe6f4dccc8d511858b9c6c30ead9ad75b0315be3adfdf55a3a91b93105cfbb22025b06c66c

 $ go run main.go -algorithm=sha512 alex
35f319ca1dfc9689f5a33631c8f93ed7c3120ee7afa05b1672c7df7b71f63a6753def5fd3ac9db2eaf90ccab6bff31a486b51c7095ff958d228102b84efd7736

 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	h "github.com/guidorice/gopl.io/ch4/ex4_2/hasher"
)

var algoPtr *string;
var helpPtr *bool;
var algo = h.Sha256

func usage() {
	flag.PrintDefaults() // prints to stderr
}

func init() {
	algoPtr = flag.String(
		"algorithm",
		"sha256",
		"hash algorithm: sha256, sha384 or sha512",
	)
	helpPtr = flag.Bool(
		"help",
		false,
		"display help",
	)
}

func checkFlags() {
	flag.Parse()

	if *helpPtr {
		usage();
		os.Exit(1)
	}

	switch *algoPtr {
	case "sha256":
		algo = h.Sha256
	case "sha384":
		algo = h.Sha384
	case "sha512":
		algo = h.Sha512
	default:
		usage()
		os.Exit(1)
	}
}

func main() {
	var input []byte
	checkFlags()
	args := flag.Args()
	if len(args) > 0 {
		input = []byte(args[0]) // TODO handle multiple words from command line
	} else {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		input = b
	}
	if len(input) == 0 {
		usage()
		os.Exit(1)
	}
	result := h.Hash(algo, input)
	fmt.Printf("%s\n", result)
}
