package main

import "fmt"

// 파이프 라인

// 채널을 통해 한 고루틴의 출력을 다른 고루틴의 입력으로 연결하는 패턴
// 큰 작업을 하나씩 단계를 나누어 처리할 때 유용하다
// https://aidanbae.github.io/code/golang-design/pipeline/screenshot.png

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter naturals 채널에 정수를 송신
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer naturals 채널에서 수신받은 값을
	// 제곱하여 squares 채널에 송신
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer squares 채널의 값은 수신하여 print
	for {
		fmt.Println(<-squares)
	}
}

// 위의 예제에서는 출력을 무한대로 하고 있다
// 그렇다면, 파이프라인을 통해 유한한 값만을 보내려면?

// 송신자가 채널에 더 이상 값이 들어가지 않게 하고, --> 내장된 close()함수로 채널을 닫는방식(닫힌 채널에 다시 송신을 하게 된다면 panic)
// 수시자가 채널이 닫힌걸 알 수 있으면 --> 채널에서 열려있는지 검사하는 방식으로(x, ok := <-naturlas)
// 채널이 열려있을때만 출력하는 식으로 조정 가능

// 개선된 Squarer
// go func() {
// 	for {
// 		x, ok := <-naturals
// 		if !ok {
// 			break
// 		}
// 		squares <- x * x
// 	}
// 	close(squares)
// }()

// 위의 문법은 맞긴 하지만, Go 언어에서는 range 루프로 close된 채널을 확인할 수 있다

// 다음의 pipeline2 예제에서는 close와 range루프를 이용해서 개선되 코드를 확인할 수 있다
