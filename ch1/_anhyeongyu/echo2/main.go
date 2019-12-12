// Echo2는 커맨드라인 인수를 출력한다.
package main

import (
	"fmt"
	"os"
)

func main() {
	// 변수 선언 및 초기화
	// s := ""
	// var s string
	// var s = ""
	// var s string = ""

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		// s의 이전 값은 더 이상 사용되지 않음 -> 가비지 컬렉션
		// 데이터의 양이 큰 경우, 많은 비용
		// 방안) strings.Join 사용
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
