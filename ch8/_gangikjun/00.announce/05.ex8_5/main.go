package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// CPU를 많이 사용하는 순차 프로그램을, 채널을 사용해 동시성 프로그래밍으로 변경
// 연습문제에서 추천하는 예제는 제가 완벽히 이해가 안되서...
// 기존 파이프라인3번 예제 수정버전

// 실패 경험 공유, 어떨때 쓰는게 이득인가?

const repeat = 10000000

// func counter() []int64 {
// 	out := make([]int64, 0, repeat)
// 	for x := int64(0); x < repeat; x++ {
// 		out = append(out, x)
// 	}
// 	return out
// }

// func squarer(in []int64) []int64 {
// 	out := make([]int64, 0, repeat)
// 	for v := range in {
// 		out = append(out, int64(v)*int64(v))
// 	}
// 	return out
// }

// func printer(in []int64) {
// 	for _, v := range in {
// 		if v%(repeat*10000000) == 0 {
// 			fmt.Println(v)
// 		}
// 	}
// }

// func main() {
// 	start := time.Now()
// 	cnt := counter()
// 	fmt.Println("counter", time.Since(start))

// 	start = time.Now()
// 	squ := squarer(cnt)
// 	fmt.Println("squarer", time.Since(start))

// 	start = time.Now()
// 	printer(squ)
// 	fmt.Println("printer", time.Since(start))
// }

//--------------------------------------------------------------------------------

// func counter(done <-chan struct{}) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for x := 0; x < repeat; x++ {
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
// 			if v%(repeat*10000000) == 0 {
// 				fmt.Println(v)
// 			}
// 		}
// 	}
// }

// func main() {
// 	start := time.Now()
// 	done := make(chan struct{})
// 	defer close(done)

// 	printer(done, squarer(done, counter(done)))
// 	fmt.Println(time.Since(start))
// }

//------------------------------------------------------------------------------------

func counter(done <-chan struct{}, s, e int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for x := s; x < e; x++ {
			select {
			case <-done:
				return
			default:
				out <- x
			}
		}
	}()
	return out
}

func squarer(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-done:
				return
			default:
				out <- v * v
			}
		}
	}()
	return out
}

func printer(done <-chan struct{}, in <-chan int) {
	for v := range in {
		select {
		case <-done:
			return
		default:
			if v%(repeat*10000000) == 0 {
				fmt.Println(v)
			}
		}
	}
}

func fanIn(done <-chan struct{}, chans ...<-chan int) <-chan int {
	out := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(chans))

	for _, c := range chans {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				select {
				case <-done:
					return
				case out <- v:
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		defer close(out)
	}()

	return out
}

func main() {
	start := time.Now()
	done := make(chan struct{})
	defer close(done)

	// numCPU := runtime.NumCPU()
	// counters := make([]<-chan int, numCPU)
	// for i := 0; i < numCPU; i++ {
	// 	startIdx := repeat / numCPU * (i)
	// 	endIdx := startIdx + repeat/numCPU
	// 	if i == numCPU-1 {
	// 		endIdx = repeat
	// 	}
	// 	counters[i] = counter(done, startIdx, endIdx)
	// }
	// countCh := fanIn(done, counters...)

	// squarers := make([]<-chan int, numCPU)
	// for i := 0; i < numCPU; i++ {
	// 	squarers[i] = squarer(done, countCh)
	// }
	// squarerCh := fanIn(done, squarers...)

	// for i := 0; i < numCPU; i++ {
	// 	printer(done, squarerCh)
	// }

	// printer(done, squarer(done, counter(done, 0, repeat)))

	numCPU := runtime.NumCPU()
	counters := make([]<-chan int, numCPU)
	for i := 0; i < numCPU; i++ {
		startIdx := repeat / numCPU * (i)
		endIdx := startIdx + repeat/numCPU
		if i == numCPU-1 {
			endIdx = repeat
		}
		counters[i] = counter(done, startIdx, endIdx)
	}
	for range fanIn(done, counters...) {
	}

	// for range counter(done, 0, repeat) {
	// }

	fmt.Println(time.Since(start))
}
