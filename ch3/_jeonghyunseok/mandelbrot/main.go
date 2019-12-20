// Mandelbrot 프랙탈 이미지 만들기
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {

	makePng("mandelbrot.png", mandelbrot)
	makePng("acos.png", acos)
	makePng("sqrt.png", sqrt)
	makePng("newton.png.png", newton)

}

func makePng(name string, dFunc func(z complex128) color.Color) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, dFunc(z))
		}
	}
	f, _ := os.Create(name)
	defer f.Close()
	png.Encode(f, img) // ignore errors
}

func mandelbrot(z complex128) color.Color {
	const iter = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iter; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

func newton(z complex128) color.Color {
	const iter = 37
	const contrast = 7
	for i := uint8(0); i < iter; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}

	}
	return color.Black
}

/*
go run main.go

*/
