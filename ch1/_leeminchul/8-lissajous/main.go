package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // color.White
	blackIndex = 1 // color.Black
)

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // x 진동자 회전 수
		res     = 0.001 // 회전 각
		size    = 100   // 캔버스 크기
		nframes = 64    // 프레임 수
		delay   = 8     // 10ms 단위, 프레임 간 딜레이
	)
	freq := rand.Float64() * 3.0 // y 진동자 상대적 진동수
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 위상차

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex) // 반올림
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // XXX: Unhandled encoding error
}
func main() {
	lissajous(os.Stdout)
}
