// Surface 는 3D 표현 함수의 SVG 렌더링을 계산한다
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 300            // 픽셀 단위 캔버스 크기
	cells         = 100                 // 격자 셀의 숫자
	xyrange       = 30.0                // 축 범위(-xyrange .. + xyrange)
	xyscale       = width / 2 / xyrange // x 나 y의 단위 픽셀
	zscale        = height * 0.4        // z 단위 픽셀
	angle         = math.Pi / 6         // x, y 축의 각도(=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30), cos(30)

func main() {
	fmt.Printf("<svg xmln='http:/www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// (i, j) 셀 코너에서 (x, y) 점 찾기
	x := xyrange * (float64(j)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 표면 높이 z 연산
	z := f(x, y)

	// (x, y, z)를 3차원 SVG 평면 (sx, sy)에 등각 투영시킴
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (0,0)에서의 거리
	return math.Sin(r)    // r
}
