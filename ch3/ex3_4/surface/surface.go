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
	Cells   = 100         // number of grid Cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

// xyScale returns pixels per x or y unit
func xyScale(width int) float64 {
	return float64(width) / 2 / xyrange
}

// zScale returns pixels per z unit
func zScale(height int) float64 {
	return float64(height) * 0.4
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// Corner returns the x,y coordinates for the corner of a cell. If the height
// calculation returns a non-finite value, an error is returned.
func Corner(fName string, width, height, i, j int) (float64, float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/Cells - 0.5)
	y := xyrange * (float64(j)/Cells - 0.5)

	// Compute surface Height z.
	z, err := f(fName, x, y)
	if err != nil {
		return 0, 0, 0, err
	}
	switch {
	case z == math.NaN() || z == math.Inf(1) || z == math.Inf(-1):
		z = 0
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyScale(width)
	sy := float64(height)/2 + (x+y)*sin30*xyScale(width) - z*zScale(height)
	return sx, sy, z, nil
}

// f computes surface Height
func f(fName string, x, y float64) (float64, error) {

	r := math.Hypot(x, y) // distance from (0,0)
	switch fName {
	case "Sin":
		return math.Sin(r) / r, nil
	case "Cos":
		return math.Cos(r) / r, nil
	case "Atan":
		return math.Atan(r) / r, nil
	case "Tan":
		return math.Tan(r) / r, nil
	case "Rand":
		return rand.Float64() / r, nil
	default:
		return 0, fmt.Errorf("surface.f(): invalid function name: %v", fName)
	}
}
