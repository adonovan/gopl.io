package chn

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// 9.5 두 버퍼되지 않은 채널 간에 주고받는 식으로 메시지를 전달하는 두 고루틴이 있는 프로그램을 작성하라.
// 이 프로그램은 초당 얼마나 많은 통신을 유지할 수 있는가?

func main() {
	var i int64
	ch1 := make(chan int)
	// ch2 := make(chan int)
	// Q. Why not working two channel?
	start := time.Now()
	go func() {
		ch1 <- 1
		for {
			i++
			ch1 <- <-ch1
		}
	}()

	go func() {
		for {
			ch1 <- <-ch1
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println(float64(i)/(float64(time.Since(start))/1e9), "round trips per sec")
}

// 2.566432802741564e+06 round trips per sec