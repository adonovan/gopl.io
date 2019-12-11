package main

import "fmt"

func main() {
	fmt.Printf("%g\n", BoilingC - FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF - CToF(FreezingC))

}