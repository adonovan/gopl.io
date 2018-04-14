/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package main

import (
	"fmt"
	"log"

	"github.com/guidorice/gopl.io/ch3/ex3_3"
	"github.com/guidorice/gopl.io/ch3/ex3_3/surface"
)

func main() {

	// first collect all the polygon points (ax, ay, bx, by, cx, cy, dx, dy)
	// and their surface heights.
	var err error

	polys := make([]ex3_3.SquarePolygon, 0)

	for i := 0; i < surface.Cells; i++ {
		for j := 0; j < surface.Cells; j++ {

			poly := ex3_3.SquarePolygon{}

			poly.Ax, poly.Ay, poly.Az, err = surface.Corner(i+1, j)
			// if the corner height calculation was non-finite, an error is
			// returned. do not emit a polygon in this case
			if err != nil {
				log.Println(err)
				continue
			}
			poly.Bx, poly.By, poly.Bz, err = surface.Corner(i, j)
			if err != nil {
				log.Println(err)
				continue
			}
			poly.Cx, poly.Cy, poly.Cz, err = surface.Corner(i, j+1)
			if err != nil {
				log.Println(err)
				continue
			}
			poly.Dx, poly.Dy, poly.Dz, err = surface.Corner(i+1, j+1)
			if err != nil {
				log.Println(err)
				continue
			}
			polys = append(polys, poly)
		}
	}
	peak := ex3_3.Peak(polys)
	peakHeight := peak.Height()
	valley := ex3_3.Valley(polys)
	valleyHeight := valley.Height()

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", surface.Width, surface.Height)

	for _, p := range polys {
		fill := ex3_3.Colorize(p.Height(), valleyHeight, peakHeight)
		fmt.Printf("<polygon fill='rgb(%v, %v, %v)' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			int(fill.R*255),
			int(fill.G*255),
			int(fill.B*255),
			p.Ax, p.Ay, p.Bx, p.By, p.Cx, p.Cy, p.Dx, p.Dy,
		)
	}

	fmt.Println("</svg>")
}
