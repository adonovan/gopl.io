package main

import "fmt"

const boillingF = 212.0

func main() {
	var f = boillingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
