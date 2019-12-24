package main

import "fmt"

func f(x int) {
	fmt.Printf("%d : ", x)
	fmt.Printf("f(%d)\n", x + 0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main() {
	f(4)
}
