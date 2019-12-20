package main

import "fmt"

var a, b, c int = 1, 2, 3
var d, e, f = 1, 2, 3

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	fmt.Println(gcd(4, 6))
}
