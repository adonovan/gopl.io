// 공백임포트를 챙겨보는 예제

package main

import (
	"fmt"
	"image"
	"image/jpeg"

	_ "image/png" // register PNG decoder
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 5})
}

/*
//!+with
go build github.com/gopl-study/gopl.io/ch3/_jeonghyunseok/mandelbrot
go build -o pkg.exe github.com/gopl-study/gopl.io/ch10/_jeonghyunseok/jpeg
mandelbrot.exe | pkg.exe >mandelbrot.jpg
Input format = png
//!-with

//!+without
go build -o nopkg.exe github.com/gopl-study/gopl.io/ch10/_jeonghyunseok/jpeg
mandelbrot | nopkg.exe >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/

// 참고
/*
https://github.com/golang/go/blob/master/src/image/png/reader.go

func init() {
	image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
}
*/
