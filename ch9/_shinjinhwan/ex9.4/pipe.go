package main

// 9.4 임의의 개수의 고루틴을 채널로 연결하는 파이프라인을 구축하라.
// 메모리 부족이 발생하기 전까지 파이프라인 단계를 최대 몇 개 까지 만들 수 있는가? recursive?
// 값이 전체 파이프라인을 거치는데 얼마나 오래 걸리는가?

// random number : 2
// tile <-> gopher <-> gopher <-> destination
//

// random number : 5
// tile <-> gopher <-> gopher <-> gopher <-> gopher <-> gopher <-> destination
//

func pipeline(gophers int) (in chan int, out chan int) {
	out = make(chan int)
	first := out
	for i := 0; i < gophers; i++ {
		in = out
		out = make(chan int)
		go func(in chan int, out chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(in, out)
	}
	return first, out
}