package geometry

import "math"

type Point struct{ X, Y float64 }

//전통적인 함수
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//동일한 역할을 수행하는 Point 타입의 메소드
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Path는 선분에서 점을 잇는 경로다.
type Path []Point

//Distance는 경로를 따라 이동한 거리를 반환한다.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
