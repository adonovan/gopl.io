// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/convertions"
	"gopl.io/ch2/tempconv"
)

func main() {

	length := flag.Bool("l", false, "Perform length conversions")
	weigth := flag.Bool("w", false, "Perform weight conversions")
	temperature := flag.Bool("t", false, "Perform temperature conversions")
	flag.Parse()

	for _, arg := range flag.Args() {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		convert := *length || *weigth || *temperature

		if !convert {
			flag.Usage()
			fmt.Println("At least one of the converions selector flag is required")
			os.Exit(1)
		}

		if *length {
			m := convertions.Meter(t)
			f := convertions.Feet(t)
			fmt.Printf("Length: %s = %s, %s = %s\n",
				m, convertions.Meter2Feet(m), f, convertions.Feet2Meter(f))
		}

		if *weigth {
			k := convertions.Kilogram(t)
			p := convertions.Pound(t)
			fmt.Printf("Weigth: %s = %s, %s = %s\n",
				k, convertions.Kilogram2Pound(k), p, convertions.Pound2Kilogram(p))
		}

		if *temperature {
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("Temperature: %s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
		}

	}
}

//!-
