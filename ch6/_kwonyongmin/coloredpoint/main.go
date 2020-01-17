package main

import (
	"fmt"
	"image/color"
	"math"
)

//Point ...
type Point struct{ X, Y float64 }

//ColoredPoint ...
type ColoredPoint struct {
	Point
	Color color.RGBA
}

//Distance ...
func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

//ScaleBy ...
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}

	fmt.Println(p.Distance(q.Point))

	p.ScaleBy(2)
	q.ScaleBy(2)

	fmt.Println(p.Distance(q.Point))
}

func init() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)

	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}

func init() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	type ColoredPoint struct {
		*Point
		color color.RGBA
	}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}

	fmt.Println(p.Distance(*q.Point))
	q.Point = p.Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point)
}
