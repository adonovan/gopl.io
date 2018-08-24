/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package main

import (
	"fmt"
	"log"

	"github.com/guidorice/gopl.io/ch3/ex3_2/surface"
)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", surface.Width, surface.Height)
	for i := 0; i < surface.Cells; i++ {
		for j := 0; j < surface.Cells; j++ {
			ax, ay, err := surface.Corner(i+1, j)
			// if the corner height calculation was non-finite, an error is
			// returned. do not emit a polygon in this case
			if err != nil {
				log.Println(err)
				continue
			}
			bx, by, err := surface.Corner(i, j)
			if err != nil {
				log.Println(err)
				continue
			}
			cx, cy, err := surface.Corner(i, j+1)
			if err != nil {
				log.Println(err)
				continue
			}
			dx, dy, err := surface.Corner(i+1, j+1)
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}
