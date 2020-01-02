package main

import "fmt"

var i = 10

func init() {
	fmt.Println("i=", i)
}

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

// 신기하다.
// 첫 fmt.Printf 도 실행되고, 이후 f(x-1) 도 실행되며,
// 이후 맨 마지막 호출된 함수부터 하나씩 defer 가 실행된다.
