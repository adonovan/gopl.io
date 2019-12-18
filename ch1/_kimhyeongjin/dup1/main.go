// Dup1은 표준입펵에서 두 번 이상 나타나는 각 줄을
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
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: input.Err() 에서의 잠재적 오류는 무시한다.
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
