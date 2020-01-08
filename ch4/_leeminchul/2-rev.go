package main

import "fmt"

func reverse(s []int) { // 파라미터로 slice 를 받는다
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // [5 4 3 2 1 0]
}
