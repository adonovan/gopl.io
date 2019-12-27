package main

import "fmt"
import "gopl.io/ch2/_beomyonglee/tempconv"

func main() {
	fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.CToF(tempconv.FreezingC))
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
