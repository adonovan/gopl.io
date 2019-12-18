package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8,8}, 5}, 20}
	w = Wheel {
		Circle : Circle{
			Point: Point{1,2},
			Radius: 4,
		},
		Spokes:20,
	}

	fmt.Printf("%#v\n", w)
	w.X = 43
	fmt.Printf("%#v\n", w)
}
