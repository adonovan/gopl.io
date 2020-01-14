package main

import "fmt"

// squares 는 호출될 때마다
// 다음번 제곱 값을 반환하는 함수를 반환한다
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
