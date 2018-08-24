// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.

// ex 1.12 modify the lissajous server to read parameter values from the url.
// for example you might arrange it so that a URL like http://localhost:8000
// /?cycles=20 sets the number of cycles to 20 instead of the default 5.
// use the strconv.Atoi function to convert the string parameter into an
// integer.

// usage notes: for best results, block favicon.ico requests in browser,
// because that will cause extraneous log output. some example query parameters:
// http://localhost:8000/?cycles=20&size=512&nframes=100&delay=1&res=0.0001

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

type config struct {
	cycles  int     // number of complete x oscillator revolutions
	res     float64 // angular resolution
	size    int     // image canvas covers [-size..+size]
	nframes int     // number of animation frames
	delay   int     // delay between frames in 10ms units
}

// defaultConfig: return a config populated with some cool defaults.
func defaultConfig() config {
	c := config{}
	c.cycles = 5
	c.res = 0.001
	c.size = 500
	c.nframes = 64
	c.delay = 8
	return c
}

// ROYGBIV colors
var palette = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // red
	color.RGBA{0xff, 0xa5, 0x00, 0xff}, // orange
	color.RGBA{0xff, 0xff, 0x00, 0xff}, // yellow
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // green
	color.RGBA{0x00, 0x00, 0xff, 0xff}, // blue
	color.RGBA{0x4b, 0x00, 0x82, 0xff}, // indigo
	color.RGBA{0xee, 0x82, 0xee, 0xff}, // violet
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout, defaultConfig())
}

func handler(w http.ResponseWriter, r *http.Request) {
	conf := defaultConfig()
	if err := r.ParseForm(); err != nil {
		log.Printf("error parsing request parameters, got %v", err)
		log.Printf("using default config: %v", conf)
		lissajous(w, conf)
		return
	}
	for k, v := range r.Form {
		switch k {
		case "cycles":
			c, err := strconv.Atoi(v[0]);
			if err != nil {
				log.Printf("parsing cycles %v, got %v", v[0], err)
			} else {
				conf.cycles = c
			}
		case "res":
			r, err := strconv.ParseFloat(v[0], 32);
			if err != nil {
				log.Printf("parsing res %v, got %v", v[0], err)
			} else {
				conf.res = r
			}
		case "size":
			s, err := strconv.Atoi(v[0]);
			if err != nil {
				log.Printf("parsing res %v, got %v", v[0], err)
			} else {
				conf.size = s
			}
		case "nframes":
			n, err := strconv.Atoi(v[0]);
			if err != nil {
				log.Printf("parsing res %v, got %v", v[0], err)
			} else {
				conf.nframes = n
			}
		case "delay":
			d, err := strconv.Atoi(v[0]);
			if err != nil {
				log.Printf("parsing res %v, got %v", v[0], err)
			} else {
				conf.delay = d
			}
		}
	}
	log.Printf("using config: %v", conf)
	lissajous(w, conf)
}

func lissajous(out io.Writer, conf config) {

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: conf.nframes}
	phase := 0.0 // phase difference
	paletteIndex := 0

	for i := 0; i < conf.nframes; i++ {
		rect := image.Rect(0, 0, 2*conf.size+1, 2*conf.size+1)
		img := image.NewPaletted(rect, palette)
		if i%2 == 0 {
			paletteIndex++
			if paletteIndex > len(palette)-1 {
				paletteIndex = 1
			}
		}
		for t := 0.0; t < float64(conf.cycles)*2*math.Pi; t += conf.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// color each few frames differently

			img.SetColorIndex(
				conf.size+int(x*float64(conf.size)+0.5),
				conf.size+int(y*float64(conf.size)+0.5),
				uint8(paletteIndex),
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, conf.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
