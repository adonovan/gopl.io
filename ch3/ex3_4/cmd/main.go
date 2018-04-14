/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	app "github.com/guidorice/gopl.io/ch3/ex3_4"
	"github.com/guidorice/gopl.io/ch3/ex3_4/surface"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	// TODO: use these query parameters, or defaults, for calls to Surface and rendering of svg
	//width := r.FormValue("width")
	//height := r.FormValue("height")
	//color1 := r.FormValue("color1")
	//color2 := r.FormValue("color2")
	//function := r.FormValue("function")

	// first collect all the polygon points (ax, ay, bx, by, cx, cy, dx, dy)
	// and their surface heights.
	var err error

	polys := make([]app.SquarePolygon, 0)

	for i := 0; i < surface.Cells; i++ {
		for j := 0; j < surface.Cells; j++ {

			poly := app.SquarePolygon{}

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
	peak := app.Peak(polys)
	peakHeight := peak.Height()
	valley := app.Valley(polys)
	valleyHeight := valley.Height()

	w.Header().Set("content-type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", surface.Width, surface.Height)

	for _, p := range polys {
		fill := app.Colorize(p.Height(), valleyHeight, peakHeight)
		fmt.Fprintf(w, "<polygon fill='rgb(%v, %v, %v)' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			int(fill.R*255),
			int(fill.G*255),
			int(fill.B*255),
			p.Ax, p.Ay, p.Bx, p.By, p.Cx, p.Cy, p.Dx, p.Dy,
		)
	}
	fmt.Fprintf(w, "</svg>")
}
