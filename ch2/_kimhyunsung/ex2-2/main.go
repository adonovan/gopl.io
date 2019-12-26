package main

import (
	"bufio"
	"fmt"
	"gopl.io/ch2/_kimhyunsung/tempconv"
	"os"
	"strconv"
)

func main() {
	var inputArg []string
	if len(os.Args) <= 1 {

		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			inputArg = append(inputArg, input.Text())
		}

		fmt.Println(inputArg)

	} else {
		for _, arg := range os.Args[1:] {
			inputArg = append(inputArg, arg)

		}
	}

	for _, arg := range inputArg {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
