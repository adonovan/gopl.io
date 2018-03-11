// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// ex 2.2 write a general purpose conversion program analogous to cf that reads
// numbers from it's command line arguments or from the standard input if there
// are no arguments, and convert each number into units like temperature in
// celcius and fahrenheit, length in feet and meters, weight in pounds and
// kilograms, and the like.

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/guidorice/gopl.io/ch2/ex2_2/conv"
)

func main() {

	var f = flag.Float64("f", math.Inf(1), "fahrenheight")
	var c = flag.Float64("c", math.Inf(1), "celsius")
	var k = flag.Float64("k", math.Inf(1), "kelvin")
	var out = flag.String("out", "", "output: c, f, or k")

	flag.Parse()

	if *out == "" {
		flag.Usage()
		os.Exit(1)
	}

	switch {
	case *f != math.Inf(1) && *out == "c":
		// FToC
		fmt.Println(conv.FToC(conv.Fahrenheit(*f)))
	case *c != math.Inf(1) && *out == "f":
		// CToF
		fmt.Println(conv.CToF(conv.Celsius(*c)))
	case *c != math.Inf(1) && *out == "k":
		// CToK
		fmt.Println(conv.CToK(conv.Celsius(*c)))
	case *k != math.Inf(1) && *out == "c":
		// KToC
		fmt.Println(conv.KToC(conv.Kelvin(*k)))
	default:
		log.Fatal("conv: sorry, unsupported conversion")
		os.Exit(1)
	}
}
