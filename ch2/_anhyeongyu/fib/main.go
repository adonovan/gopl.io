// 피보나치 수열
package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib(10))
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		// fmt.Printf("%d %d\n", x, y)
	}
	return x
}
