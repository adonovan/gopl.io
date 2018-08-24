/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package ex3_4

import "math"

// Valley finds the square polygon with min height by averaging the corner heights
func Valley(polys []SquarePolygon) SquarePolygon {
	min := math.MaxFloat64
	var minIdx int

	for i, p := range polys {
		z := p.Height()
		if z < min {
			min = z
			minIdx = i
		}
	}
	return polys[minIdx]
}
