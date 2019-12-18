// Lissajous는 임의의 리사주 형태의 애니메이션 GIF를 생성한다.
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
	whiteIndex = 0 // 팔레트의 첫 번째 색상
	blackIndex = 1 // 팔레트의 다음 색상
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5 // x 진동자의 회전수
		res = 0.001 // 회전각
		size = 100 // 이미지 캔버스 크기 [-size..+size]
		nframes = 64 // 애니메이션 프레임 수
		delay = 8 // 10ms 단위의 프레임 간 지연
	)
	freq := rand.Float64() * 3.0 // y 진동자의 상대적 진동수
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 위상 차이
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: 인코딩 오류 무시
}