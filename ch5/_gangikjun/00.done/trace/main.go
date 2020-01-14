package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()

	// ... 다수의 작업 ...

	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
