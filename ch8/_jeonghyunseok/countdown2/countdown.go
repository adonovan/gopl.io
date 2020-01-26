// 이건 1과 뭐가 다를까?
// 1을 복습해보면 아래와 같다.package countdown2

/*
tick := time.Tick(1*time.Second) 함수를 하나 만들고
<- tick 이 되도록 했음. 요건 time 을 주고받는 chan 임
*/

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan int)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- 0
	}()

	tick := time.Tick(1 * time.Second)

	fmt.Println("Commencing countdown. Press return to abort.")
	for i := 10; i > 0; i-- {
		select {
		case <-tick:
			// do nothing
			fmt.Println("countdown: ", i)

		case <-abort:
			fmt.Println("launch abort!")
			return
		}
	}
	time.Sleep(1*time.Second)
	launch()
}

// func main() {
// 	abort := make(chan int)
// 	go func() {
// 		os.Stdin.Read(make([]byte, 1))
// 		abort <- 0
// 	}()

// 	fmt.Println("Commencing countdown. Press return to abort.")
// 	select {
// 	case <-time.After(10 * time.Second):
// 		// do nothing
// 	case <-abort:
// 		fmt.Println("launch abort!")
// 		return
// 	}
// 	launch()
// }

func launch() {
	fmt.Println("Lift off!")
}
