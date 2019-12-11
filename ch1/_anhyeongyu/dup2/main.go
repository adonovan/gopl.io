// Dup2는 입력에서 두 번 이상 나타나는 각 줄의 카운트와 텍스트를 출력한다.
// 이 프로그램은 표준 입력이나 파일 목록에서 읽는다.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// os.Args : Command Line 의 Arguments 를 얻기 위해 사용
	// os.Args[0:1] -> 실행되는 Go 프로그램의 이름
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				// 파일이 없는 경우, 아래 과정 필요 없기 때문
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: input.Err()에서의 잠재적 오류는 무시한다.
}
