package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

// 함수
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 메서드
// 경로를 따라 이동한 거리를 return
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Path 는 선분에서 점을 잇는 경로다
type Path []Point

// 경로를 따라 이동한 거리
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // 5
	fmt.Println(p.Distance(q))  // 5

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // 12
}
