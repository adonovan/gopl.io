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
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cfg := NewLissajousConfig(r.URL.String())
		lissajous(w, cfg)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var palette = []color.Color{color.White, color.Black}

const (
	WhiteIndex = 0 // palette color index -0
	BlackIndex = 1 // palette color index -1
)

type LissajousConfig struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func NewLissajousConfig(rURL string) LissajousConfig {
	cfg := LissajousConfig{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}
	u, err := url.Parse(rURL)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	for k, _ := range q {
		v := q.Get(k)
		if v == "" {
			continue
		}

		if k == "cycles" {
			cfg.cycles, _ = strconv.Atoi(v)
		} else if k == "size" {
			cfg.size, _ = strconv.Atoi(v)
		} else if k == "nframe" {
			cfg.nframes, _ = strconv.Atoi(v)
		} else if k == "delay" {
			cfg.delay, _ = strconv.Atoi(v)
		} else if k == "res" {
			cfg.res, _ = strconv.ParseFloat(v, 64)
		}
	}
	return cfg
}

func lissajous(out io.Writer, cfg LissajousConfig) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: cfg.nframes}
	phase := 0.0
	for i := 0; i < cfg.nframes; i++ {
		rect := image.Rect(0, 0, 2*cfg.size+1, 2*cfg.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cfg.cycles)*2*math.Pi; t += cfg.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(cfg.size+int(x*float64(cfg.size)+0.5),
				cfg.size+int(y*float64(cfg.size)+0.5), BlackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, cfg.delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
