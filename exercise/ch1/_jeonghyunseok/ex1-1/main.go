// 1.1 echo 프로그램을 수정해 호출한 명령인 os.Args[0]도 같이 출력하라.
// start from Echo3

// Ex1-1 는 명령어까지 함께 보여준다.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[0], ":", strings.Join(os.Args[1:], " "))
}

/*
go build -o ex1-1.exe
ex1-1.exe Who let the dogs out!
*/
