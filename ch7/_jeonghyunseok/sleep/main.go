// 일정시간 동안 잠드는 기능

package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	for i := range []int{1, 2, 3, 4, 5} {
		_ = i
		fmt.Printf("Type: %T, Sleeping for %v...", period, *period)
		time.Sleep(*period)
		fmt.Printf("wake up: %v\n", time.Now().Second())
	}
}

/*
go build -o sleep.exe
sleep.exe -period 2s

*/
