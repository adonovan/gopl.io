/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package main

// TODO use gradient colors from http options

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	app "github.com/guidorice/gopl.io/ch3/ex3_4"
	"github.com/guidorice/gopl.io/ch3/ex3_4/surface"
)

// opts is a struct for the request parameters, or use defaults.
type opts struct {
	width     int
	height    int
	hexColor1 string
	hexColor2 string
	function  string
}

// getOpts takes for formValue()s from the http.Request r, or applies defaults.
func getOpts(r *http.Request) opts {
	opts := opts{
		width:     800,
		height:    600,
		hexColor1: "#9e0142",
		hexColor2: "#5e4fa2",
		function:  "Sin",
	}
	width := r.FormValue("width")
	var err error
	if len(width) > 0 {
		opts.width, err = strconv.Atoi(width)
		if err != nil {
			log.Fatalf("getOpt: %v", err)
		}
	}
	height := r.FormValue("height")
	if len(height) > 0 {
		opts.height, err = strconv.Atoi(height)
		if err != nil {
			log.Fatalf("getOpt: %v", err)
		}
	}
	color1 := r.FormValue("color1")
	if len(color1) > 0 {
		opts.hexColor1 = color1
	}
	color2 := r.FormValue("color2")
	if len(color2) > 0 {
		opts.hexColor2 = color2
	}
	function := r.FormValue("function")
	if len(function) > 0 {
		opts.function = function
	}
	return opts
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	opts := getOpts(r)
	log.Printf("%v", opts)

	// first collect all the polygon points (ax, ay, bx, by, cx, cy, dx, dy)
	// and their surface heights.
	var err error

	polys := make([]app.SquarePolygon, 0)

	for i := 0; i < surface.Cells; i++ {
		for j := 0; j < surface.Cells; j++ {

			poly := app.SquarePolygon{}

			poly.Ax, poly.Ay, poly.Az, err = surface.Corner(
				opts.function,
				opts.width,
				opts.height,
				i+1,
				j,
			)
			// if the corner height calculation was non-finite, an error is
			// returned. do not emit a polygon in this case
			if err != nil {
				http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
				return
			}
			poly.Bx, poly.By, poly.Bz, err = surface.Corner(
				opts.function,
				opts.width,
				opts.height,
				i,
				j,
			)
			if err != nil {
				http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
				return
			}
			poly.Cx, poly.Cy, poly.Cz, err = surface.Corner(
				opts.function,
				opts.width,
				opts.height,
				i,
				j+1,
			)
			if err != nil {
				http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
				return
			}
			poly.Dx, poly.Dy, poly.Dz, err = surface.Corner(
				opts.function,
				opts.width,
				opts.height,
				i+1,
				j+1,
			)
			if err != nil {
				http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
				return
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
		"width='%d' height='%d'>", opts.width, opts.height)

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
