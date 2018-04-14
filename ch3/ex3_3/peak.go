/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package ex3_3

// Peak finds the square polygon with peak height by averaging the corner heights
func Peak(polys []SquarePolygon) SquarePolygon {
	var max float64
	var maxIdx int

	for i, p := range polys {
		z := p.Height()
		if z > max {
			max = z
			maxIdx = i
		}
	}
	return polys[maxIdx]
}
