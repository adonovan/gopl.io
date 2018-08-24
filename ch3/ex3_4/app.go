/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package ex3_4

// package types

type SquarePolygon struct {
	// corner x,y coordinates
	Ax,
	Ay,
	Bx,
	By,
	Cx,
	Cy,
	Dx,
	Dy,
	// corner surface heights
	Az,
	Bz,
	Cz,
	Dz float64
}

// Height averages the heights Az, Bz, Cz, Dz
func (sp SquarePolygon) Height() float64 {
	return (sp.Az + sp.Bz + sp.Cz + sp.Dz) / 4
}
