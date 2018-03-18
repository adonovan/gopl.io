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

func main() {
	conv.TempFlags()
	conv.DistFlags()
	conv.WeightFlags()
	toScale := flag.String("out", "", "output scale: c, f, or k")
	flag.Parse()
	if *toScale == "" {
		flag.Usage()
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
	log.Fatal("conv: sorry, unsupported conversion")
}
