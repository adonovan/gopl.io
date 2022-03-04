package main

import (
	"fmt"
)

func main()  {
	x:= make(map[string]int)
	y:=make(map[string]int)
	x["3bu1"] = 31
	x["sneha"] = 32
	x["archana"] = 0
	y["3bu1"] = 31
	y["sneha"] = 32
	y["shivani"] = 0
	z :=equals(x,y)
	fmt.Println("z : ",z)
			fmt.Println("dd")
}

func equals(x,y map[string]int) bool  {
	fmt.Println(x,y)
	if len(x) != len(y) {
		return false
	}
	for xk, xv := range x {
		if yv, ok:= y[xk]; !ok || xv != yv {
			return false
		}
	}
	return true
}