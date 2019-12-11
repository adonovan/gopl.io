package main

import (
	"fmt"
	"os"
	"strconv"
	"../tempconv0"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv0.Fahrenheit(t)
		c := tempconv0.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv0.FToC(f), c, tempconv0.CToF(c))

	}
}
