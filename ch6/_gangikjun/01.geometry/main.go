package main

import (
	"bytes"
	"fmt"
)

// type ptr *int

// func (p ptr) f()  {}
// func (p *ptr) f() {}

// type in interface{}

// func (i in) f()  {}
// func (i *in) f() {}

func main() {
	// p := geometry.Point{1, 2}
	// q := geometry.Point{4, 6}
	// fmt.Println(geometry.Point.Distance(p, q))
	// fmt.Println(geometry.Distance(p, q))
	// fmt.Println(p.Distance(q))

	// perim := geometry.Path{
	// 	{1, 1},
	// 	{5, 1},
	// 	{5, 4},
	// 	{1, 1},
	// }
	// fmt.Println(perim.Distance())

	// r := geometry.Point{1, 2}
	// r.ScaleBy(2)
	// fmt.Println(r)

	var buf bytes.Buffer
	buf.WriteRune('A')
	buf2 := buf
	buf2.WriteRune('a')

	//fmt.Println(buf.String())
	fmt.Println(buf2.String())
}
