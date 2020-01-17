package geometry

import (
	"testing"
)

func TestPointDistance(t *testing.T) {
	p := Point{5, 3}
	q := Point{2, 1}
	d1, d2 := Distance(p, q), p.Distance(q)
	if d1 != d2 {
		t.Errorf("Distance(%v, %v) = %f, but %v.Distance(%v) = %f\n",
			p, q, d1, p, q, d2)
	}
}

func TestPathDistance(t *testing.T) {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	expected := 12.0
	distance := perim.Distance()
	if expected != distance {
		t.Errorf("expected %f, was %f\n", expected, distance)
	}
}
