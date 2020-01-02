package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Println("n =", n)
	os.Stdout.Write(buf[:n])
	fmt.Println("printStack end")
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // 아하 0을 0으로 나눌때만 패닉이구나
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
