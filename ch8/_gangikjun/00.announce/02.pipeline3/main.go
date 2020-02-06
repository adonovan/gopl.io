package main

import "fmt"

// 단방향 채널 타입

// 채널의 오용(수신을 위한 채널에 송신을 한다거나 등등..)을 막기 위해, 단방향 채널 타입 제공
// chan <- int 타입은 int를 송신만 하는 채널로 수신 불가
// <-chan int 타입은 int를 수신만 하는 채널로 송신 불가
// 위의 규칙 위반시 컴파일 시 감지.

// close 동작은 채널에 더 이상의 송신이 일어나지 않을 것을 가정하기 때문에
// 수신 전용 채널(<-chan int 같은) 에서 닫으려고 하면 컴파일 오류

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals) // 여기서 chan int 형을 chan<- int 형으로 묵시적으로 변환
	// 양방향 채널(chan int)에서 단방향 채널로의 변환은 ok, 역은 불가
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

// 1. 채널을 생성한 고루틴이, 해당 채널을 관리하여야 한다
// 2. 고루틴은 종료시점을 명확하게 하여야 한다
// func main() {
// 	done := make(chan struct{})
// 	defer close(done)

// 	conuterCh := counter(done)
// 	squarerCh := squarer(done, conuterCh)
// 	printer(done, squarerCh)

// 	// printer(done, squarer(done, counter(done)))
// }

// func counter(done <-chan struct{}) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for x := 0; x < 100; x++ {
// 			select {
// 			case <-done:
// 				return
// 			default:
// 				out <- x
// 			}
// 		}
// 	}()
// 	return out
// }

// func squarer(done <-chan struct{}, in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for v := range in {
// 			select {
// 			case <-done:
// 				return
// 			default:
// 				out <- v * v
// 			}
// 		}
// 	}()
// 	return out
// }

// func printer(done <-chan struct{}, in <-chan int) {
// 	for v := range in {
// 		select {
// 		case <-done:
// 			return
// 		default:
// 			fmt.Println(v)
// 		}
// 	}
// }
