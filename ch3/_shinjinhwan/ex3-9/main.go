package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	const width, height          = 1024, 1024

	handler := func(w http.ResponseWriter, r *http.Request) {
		var xmin, ymin, xmax, ymax, zoom float64 = -2, -2, +2, +2, 1

		s :=  r.FormValue("zoom")
		if s != "" {
			var err error
			zoom, err = strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Fprintf(w, "bad zoom param : %f", zoom, err)
				return
			} else {
				fmt.Printf("Parsed Arguments : %f", zoom)
			}
		}

		lenX := xmax - xmin
		midX := xmin + lenX/2
		xmin = midX - lenX/2/zoom
		xmax = midX + lenX/2/zoom

		lenY := ymax - ymin
		midY := ymin + lenY/2
		ymin = midY - lenY/2/zoom
		ymax = midY + lenY/2/zoom

		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)

				img.Set(px, py, mandelbrot(z))
			}
		}
		err := png.Encode(w, img)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
