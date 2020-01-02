package main

import (
	"fmt"
)

func main() {
	type point struct {
		x, y, z int
		id      int32
		name    string
	}

	p := point{1, 2, 3, 4, "a"}
	fmt.Println(p)
}
