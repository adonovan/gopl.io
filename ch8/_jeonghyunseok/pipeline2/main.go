// 파이프라인2는 유한한 파이프라인 구현을 보여준다

package main

import "fmt"

func main() {
	nt := make(chan int)
	squares := make(chan int)

	// 카운터
	go func() {
		for x := 0; x < 100; x++ {
			nt <- x
		}
		close(nt)
	}()

	// 제곱
	go func() {
		for x := range nt {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}

}
