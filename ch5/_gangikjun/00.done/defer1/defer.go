package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // x == 0 이면 패닉
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
