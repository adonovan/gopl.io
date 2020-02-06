package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 100; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
