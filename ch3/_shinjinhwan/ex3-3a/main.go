package main

import (
	"fmt"
	"log"
	"math"
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
var z_max, z_min float64


func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	var points = [cells+1][cells+1][10]float64{}
	minmax()

	for i := 0; i < cells; i++ {
		for j := 0; j <= cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			color_code, color_num := color(i, j)
			points[i][j] = [10]float64{color_code, color_num, ax, ay, bx, by, cx, cy, dx, dy}
		}
	}

	for _, point_x := range points {
		for _, point := range point_x {
			var color_hex string
			if point[0] == 0 {
				color_hex = fmt.Sprintf("#%x0000", int(point[1]))
			}

			if point[0] == 1 {
				color_hex = ""
			}

			if point[0] == 2 {
				color_hex = fmt.Sprintf("#0000%x", int(point[1]))
			}

			log.Println(color_hex)
			fmt.Printf("<polygon stroke='%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color_hex, point[2], point[3], point[4], point[5], point[6], point[7], point[8], point[9])
		}
	}

	fmt.Println("</svg>")
	// h := fmt.Sprintf("%x", int(254 - z_min / (z_max - z_min)))
	// log.Println(h)
}

func minmax() {
	for i := 0; i < cells+1; i++ {
		for j := 0; j < cells+1; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			z := f(x, y)

			if math.IsNaN(z_min) || z < z_min {
				z_min = z
			}
			if math.IsNaN(z_max) || z > z_max {
				z_max = z
			}
		}
	}
}

func color(i, j int) (float64, float64) {
	// TODO: How to chosse mid number? median?
	// mid := z_max - ((math.Abs(z_max) + math.Abs(z_min)) / 2)
	mid := 0.5

	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	intensity := ((z - z_min) / (z_max - z_min)) * 255
	if intensity > 255 {
		intensity = 255
	}

	// above mid is red
	if z > mid {
		return 0, intensity
	}
	// below mid is blue
	if z < mid * 0.05 {
		return 2, intensity
	}

	return 1, 0
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
