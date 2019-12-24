package main

import (
	"exam01/tempconv"
	"fmt"
)

func main() {
	var k tempconv.Kelvin = 0
	fmt.Println(k)

	k++
	fmt.Println(k)

	fmt.Println(tempconv.KToC(k))
}
