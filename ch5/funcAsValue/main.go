package main

import (
	"fmt"
)

func main() {
	f := square
	fmt.Println(f(3))

	f = negative
	fmt.Println(f(3))

	f = product
}

func square(n int) int {
	return n * n
}

func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
