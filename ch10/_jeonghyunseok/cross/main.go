// GOOS, GOARCH 를 설정해서 크로스 컴파일링 할 수 있다.
package main

import "fmt"
import "runtime"

func main() {
	fmt.Printf("GOOS: %v, GOARCH: %v", runtime.GOOS, runtime.GOARCH)
}

/*
go build -o nocross.exe

set GOOS=linux
set GOARCH=amd64
go build -o cross.exe

*/
