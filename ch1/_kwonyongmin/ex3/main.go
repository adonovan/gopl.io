package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
	res := time.Now().Sub(start)
	fmt.Println("잠재적으로 비효율적인 버전 시간 :", res)

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	res = time.Now().Sub(start)
	fmt.Println("strings.Join을 사용한 버전 시간 :", res)
}
