/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// Surface computes an SVG rendering of a 3-D surface function.
package surface

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	Width, Height = 600, 320            // canvas size in pixels
	Cells         = 100                 // number of grid Cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = Width / 2 / xyrange // pixels per x or y unit
	zscale        = Height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// Corner returns the x,y coordinates for the corner of a cell. If the height
// calculation returns a non-finite value, an error is returned.
func Corner(i, j int) (float64, float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/Cells - 0.5)
	y := xyrange * (float64(j)/Cells - 0.5)

	// Compute surface Height z.
	z := f(x, y)
	switch {
	case z == math.NaN() || z == math.Inf(1) || z == math.Inf(-1):
		return 0, 0, 0, fmt.Errorf(
			"corner: surface height non-finite for i, j %v, %v, got %v", i, j, z,
		)
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := Width/2 + (x-y)*cos30*xyscale
	sy := Height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

// f computes surface Height
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	//s := math.Atan(r) / r
	//s := math.Cos(r) / r
	// s := math.Sin(r) / r
	s := r / 2 * rand.Float64() / math.Pow(r, 2)
	return s
}
