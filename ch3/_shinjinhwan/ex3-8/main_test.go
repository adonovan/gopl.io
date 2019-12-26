package main

import (
	"image/color"
	"testing"
)

func benchmarkMandelbrot(b *testing.B, f func(complex128) color.Color) {
	for i := 0; i < b.N; i++ {
		f(complex(float64(i), float64(i)))
	}
}

func BenchmarkMandelbrotComplex128(b *testing.B) {
	benchmarkMandelbrot(b, mandelbrot)
}

func BenchmarkMandelbrotBigFloat(b *testing.B) {
	benchmarkMandelbrot(b, mandelbrotBigFloat)
}

/*
goos: darwin
goarch: amd64
BenchmarkMandelbrotComplex128-8         51015427                23.7 ns/op             1 B/op          1 allocs/op
BenchmarkMandelbrotBigFloat-8            2273628               575 ns/op             232 B/op          9 allocs/op
PASS
ok      _/Users/jinhwan/gitrepo/gopl.io/ch3/_shinjinhwan/ex3-8  4.073s
 */