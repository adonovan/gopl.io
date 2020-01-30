// 파이프라인 1은 무한 3단계 파이프라인을 보여준다.
package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 카운터
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
	}()

	// 제곱
	go func() {
		for {

			x := <-naturals
			squares <- x * x
		}
	}()

	// 출력
	for {
		fmt.Println("squares: ", <-squares)
	}

}
