// 1.12 URL에서 파라미터 값을 읽도록 리사주 서버를 수정하라.
// 예를 들어 http://localhost:8000/?cycles=20과 같은 URL이
// 반복 횟수로 기본 값 5 대신에 20을 지정하게 할 수 있을 것이다.
// strconv.Atoi 함수를 사용해 문자열 파라미터를 정수로 변환하라.
// go doc strconv.Atoi로 관련 문서를 볼 수 있다."
// start with lissajous.go

// Lissajous gogo
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// web or os.Stdout
	if len(os.Args) > 1 && os.Args[1] == "web" {

		http.HandleFunc("/", handler)
		http.HandleFunc("/help/", helpHandler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout, 5, 100, 8)
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
	help := "you can send query: cycles, size, delay"
	help2 := "ex)http://localhost:8000/?cycles=20&size=200&delay=4"

	fmt.Fprint(w, help)
	fmt.Fprint(w, help2)
}

// https://blog.charmes.net/post/query-string-parsing-go/
func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %s", err)
		return
	}
	c := r.Form.Get("cycles")
	cycles, err := strconv.Atoi(c)
	if err != nil {
		log.Printf("Error parsing cycles: %s", err)
		return
	}

	s := r.Form.Get("size")
	size, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Error parsing size: %s", err)
		return
	}

	d := r.Form.Get("delay")
	delay, err := strconv.Atoi(d)
	if err != nil {
		log.Printf("Error parsing delay: %s", err)
		return
	}

	lissajous(w, cycles, size, delay)
}

func lissajous(out io.Writer, c, s, d int) {
	var (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	cycles = c
	size = s
	delay = d

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

/*
go build -o ex1-12.exe
ex1-12.exe web

on browser
http://localhost:8000/?cycles=20&size=200&delay=4
http://localhost:8000/?cycles=50&size=300&delay=8
http://localhost:8000/?cycles=50&size=300&delay=2
*/
