package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	// 100개의 정수 송신 후 close
	go func() {
		defer close(naturals)
		for x := 0; x < 100; x++ {
			naturals <- x
		}
	}()

	// Squarer
	// range 루프를 이용해 naturals의 close 확인
	go func() {
		defer close(squares)
		for x := range naturals {
			squares <- x * x
		}
	}()

	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}

// 특이사항

// 항상 다 쓴 채널을 닫을 필요는 없다
// 가비지 콜렉터는 닫혀있는지 여부와 상관 없이, 참조할 수 없는 채널의 자원을 회수한다
// (열린 파일 닫기와는 다름, 다 쓴 파일은 항상 Close호출 필요)

// 닫힌 채널 or nil 채널을 닫으려 하면 panic
