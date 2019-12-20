package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	wg.Add(2)

	// 문자열을 합치는 작업을 할때
	// normalJoin과 stringsJoin 함수를 동일하게 테스트 하기위해
	// testSlice를 두었다
	testSlice := make([]string, 10000)
	for i := 0; i < len(testSlice); i++ {
		testSlice[i] = "a"
	}

	normalJoin := func() {
		defer wg.Done()
		var s, sep string
		for _, args := range testSlice {
			s += sep + args
			sep = " "
		}
		fmt.Println("normalJoin since time :", time.Since(start))
	}

	stringsJoin := func() {
		defer wg.Done()
		strings.Join(testSlice, " ")
		fmt.Println("stringsJoin since time :", time.Since(start))
	}

	go normalJoin()
	go stringsJoin()

	wg.Wait()
}
