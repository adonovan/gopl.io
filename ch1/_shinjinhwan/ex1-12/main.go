package main

import (
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

const whiteIndex = 0
const blackIndex = 1

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			s := r.URL.Query().Get("cycles")
			cycles, err := strconv.Atoi(s)
			if err != nil {
				log.Printf("Error parsing cycle : %s", err)
				return
			} else {
				log.Printf("Parsed argument : %d", cycles)
			}
			lissajous(w, float64(cycles))
		}
		http.HandleFunc("/", handler)

		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout, 5)
}

func lissajous(out io.Writer, cycle float64) {
	const (
		res    = 0.001
		size   = 100
		nframe = 64
		delay  = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframe}
	phase := 0.0
	for i := 0; i < nframe; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
