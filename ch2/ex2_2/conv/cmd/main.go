/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/guidorice/gopl.io/ch2/ex2_2/conv"
)

const usage = `
converts weights, distances and temperatures. supported conversions: 

* weight *
pounds    -> kilograms
stone     -> kilograms
kilograms -> pounds
kilograms -> stone

* distance *
feet   -> meters
yards  -> meters
meters -> feet
meters -> yards

* temperature *
celsius -> fahrenheit
fahrenheit -> celsius
kelvin -> celsius
celsius -> kelvin

example usage (kelvin to celsius):

	conv -k 1000 -out c
	726.85°C

`

func main() {
	conv.TempFlags()
	conv.DistFlags()
	conv.WeightFlags()
	toScale := flag.String(
		"out",
		"",
		"output scale: c, f, k, ft, m, y, kg, lb, st",
	)
	flag.Parse()
	if *toScale == "" {
		flag.Usage()
		fmt.Print(usage)
		os.Exit(1)
	}
	// try temperature conversions
	temp, err := conv.TempConv(toScale)
	if err == nil {
		fmt.Println(temp)
		os.Exit(0)
	}
	// try distance conversions
	dist, err := conv.DistConv(toScale)
	if err == nil {
		fmt.Println(dist)
		os.Exit(0)
	}
	// try weight conversions
	weight, err := conv.WeightConv(toScale)
	if err == nil {
		fmt.Println(weight)
		os.Exit(0)
	}
	flag.Usage()
	fmt.Print(usage)
	log.Fatal("conv: sorry, unsupported conversion")
}
