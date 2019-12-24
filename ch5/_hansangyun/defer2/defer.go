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

func f(x int) {
	//fmt.Printf("%d : ", x)
	fmt.Printf("f(%d)\n", x + 0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	fmt.Println("This is printStack()")
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
	fmt.Println("END =========================")
}