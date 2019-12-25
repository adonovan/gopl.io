package main

import "fmt"

func main() {

	// x := 1
	// p := &x
	// fmt.Println(p)

	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F" = 0°C"
	fmt.Printf("%g°F = %g°C\n",
		boilingF, fToC(boilingF)) // "32°F" = 0°C"
}

func fToC(f float64) float64 {
	a := 3
	fmt.Println(a)
	return (f - 32) * 5 / 9
}
