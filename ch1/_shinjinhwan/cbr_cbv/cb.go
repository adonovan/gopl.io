package main

import (
	"fmt"
)

func main() {
	var a int
	var b = make([]int, 2)

	a = 2

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(&b)
}