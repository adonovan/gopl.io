package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 2014
)

func main() {
	workers := runtime.GOMAXPROCS(-1)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	start := time.Now()
	wg := sync.WaitGroup{}
	rows := make(chan int, height)
	for row := 0; row < height; row++ {
		rows <- row
	}
	close(rows)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for py := range rows {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					img.Set(px, py, newton(z))
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("rendered in:", time.Since(start))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newton(z complex128) color.Color {
	iterations := 37
	for n := uint8(0); int(n) < iterations; n++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(cmplx.Pow(z, 4)-1) < 1e-6 {
			return color.Gray{255 - uint8(math.Log(float64(n))/math.Log(float64(iterations+0))*255)}
		}
	}
	return color.Black
}

/*
실행 후
웹 주소창에 localhost:8080 입력 후 사진을 확인한다.
*/
