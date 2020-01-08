package main

import "fmt"

func reverse(s *[5]int) { // 가변 길이의 array 를 받을 방법이 없다
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // [5 4 3 2 1 0]
}
