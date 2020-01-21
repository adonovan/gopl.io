package coloredpoint

import (
	"image/color"
	"math"
)

// Point Point
type Point struct{ X, Y float64 }

// ColoredPoint ColoredPoint
type ColoredPoint struct {
	Point
	Color color.RGBA
}

// Distance distance
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy ScaleBy
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Path 는 선분에서 점을 잇는 경로다
type Path []Point

// Distance distance
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
