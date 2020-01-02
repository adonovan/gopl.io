package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}

	rotate(s)
	fmt.Println(s)

	rotate(s)
	fmt.Println(s)

	rotate(s)
	fmt.Println(s)
}

func rotate(s []int) {
	first := s[0]
	copy(s, s[1:])
	s[len(s)-1] = first
}
