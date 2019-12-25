// Dup1은 표준 입력에서 두 번 이상 나타나는 각 줄을
// 앞에 카운트를 추가해 출력한다.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// input.Scan()을 호출 할 때 마다 다음 줄을 읽고 맨 끝의 개행문자("\n")를 제거
	for input.Scan() {
		// for 문 중단
		// mac: control + d
		// windows: control + z
		counts[input.Text()]++
	}
	// NOTE: input.Err()에서의 잠재적 오류는 무시한다.
	for line, n := range counts {
		// Map은 순서가 없다!!
		// HashTable을 사용하기 때문 랜덤하게 출력
		if n > 1 {
			// f로 끝나는 포매팅 함수는 포매팅 규칙을 사용
			// %d: 10진 정수, %s: 문자열
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	// D 또는 Z가 출력될 수 있음
	// 입력 종료 때 입력 된 키가 겹쳐서 보이는 것
}
