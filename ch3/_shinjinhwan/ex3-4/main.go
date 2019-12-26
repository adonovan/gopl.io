package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	var svgfile []string

	svgfile = append(svgfile, fmt.Sprintf(`<svg xmlns='http://www.w3.org/2000/svg' ` +
		`style='stroke: grey; fill: white; stroke-width: 0.7' ` +
		`width='%d' height='%d'>`, width, height))

	for i := 0; i < cells; i++ {
		for j := 0; j <= cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			svgfile = append(svgfile, fmt.Sprintf(`<polygon points='%g,%g %g,%g %g,%g %g,%g'/>`, ax, ay, bx, by, cx, cy, dx, dy))
		}
	}
	svgfile = append(svgfile, "</svg>")
	// fmt.Print(strings.Join(svgfile, ""))
	htmltemplate := `
<!DOCTYPE>
<html>
<body>
    <h1>Lissajous</h1>
    %s
</body>
</html>`
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, fmt.Sprintf(htmltemplate, strings.Join(svgfile, "")))
	}
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
