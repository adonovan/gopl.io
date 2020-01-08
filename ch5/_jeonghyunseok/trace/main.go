// Trace 프로그램이 defer 를 이용해서 함수에 진입하고 나가는 것을 분석한다.

package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}

/*
go build -o trace.exe
trace.exe

*/

// d:\golang\src\gopl.io\ch5\_jeonghyunseok\trace>trace.exe
// 2020/01/08 18:17:28 enter bigSlowOperation
// 2020/01/08 18:17:38 exit bigSlowOperation (10.0644745s)
