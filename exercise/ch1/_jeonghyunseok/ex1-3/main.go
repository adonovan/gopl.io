/* 1.3 잠재적으로 비효율적인 버전과 strings.Join을 사용하는 버전의
  실행 시간 차이를 실험을 통해 측정하라(1.6절에서 time 패키지를 소개하며,
	11.4절은 체계적인 성능 평가를 위한 벤치마크 테스트를 작성하는 방법을 보여준다).
*/
// Use time package this time
// compare echo2, echo3
// reference: https://stackoverflow.com/questions/28755757/time-since-golang-nanosecond-timestamp

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo2(os.Args[1:])
	fmt.Println()
	echo3(os.Args[1:])

}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string) {
	fmt.Println(strings.Join(args, " "))
}

// func echo2(args []string) {
// 	start := time.Now()
// 	s, sep := "", ""
// 	for _, arg := range args {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// 	fmt.Printf("elapsed time. echo2: %d ns", time.Since(start))
// }

// func echo3(args []string) {
// 	start2 := time.Now()
// 	fmt.Println(strings.Join(args, " "))
// 	fmt.Printf("elapsed time. echo3: %d ns", time.Since(start2))
// }

/*
go build -o ex1-3.exe
ex1-3.exe Who let the dogs out!
-> this example not works properly
-> may need benchmark
*/
