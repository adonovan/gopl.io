package main

import (
	"embed/wheel"
	"fmt"
)

func main() {
	var w wheel.Wheel
	w.X = 41

	// w := wheel.Wheel{wheel.Circle{wheel.Point{8, 8}, 5}, 20}
	// w = wheel.Wheel{
	// 	Circle: wheel.Circle{
	// 		Point:  wheel.Point{X: 8, Y: 8},
	// 		Radius: 5,
	// 	},
	// 	Spokes: 20, // NOTE : 따라오는 콤마가 필요하다.(Radius에도)
	// }

	fmt.Printf("%#v\n", w)

	w.X = 42
	fmt.Printf("%#v\n", w)
}
