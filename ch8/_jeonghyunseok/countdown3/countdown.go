// 고루틴 릭이 있다고 한다. 이걸 어찌 막을까?
// 오, 이건 내가 countdown 2를 수정한것과 비슷하다.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	//!+
	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	close(abort)  // 이렇게 넣어주면 되지 않을까?
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
