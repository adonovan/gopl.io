package main

import (
	"fmt"

	"gopl.io/ch4/rev/reverse"
)

func main()  {
	a := []int{1,2,3,4,5}
	reverse.Reverse(a[:3])
	reverse.Reverse(a[3:])
	fmt.Println(a[:])
}

func equal(x []string , y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i, _ := range x {
		if x[i] != y[i] {
		   return false	
		}
	}
	return true
}

