// Colordpoint 는 구조체 임베딩을 보여준다.
package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

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
	fmt.Println("init1")
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance // 메쏘드를 가져갔네. 호호 이렇게도 되는구나, 첫째 타입이 첫째 파라미터
	fmt.Println(distance(p, q))
	fmt.Printf("type of distance: %T\n", distance)

	scale := (*Point).ScaleBy // 이건 Point 가 아니라 (*Point) 타입의 메쏘드이다!
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("type of scale: %T\n", scale)
}

func init() {
	fmt.Println("init2")
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	type ColoredPoint struct { // init 안에서의 타입 선언과 패키지 레빌 선언과의 충돌 확인
		*Point
		Color color.RGBA
	}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point))
	q.Point = p.Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point)
	fmt.Println(p.Point, q.Point)
}
