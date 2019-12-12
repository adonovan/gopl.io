// dup2는 입력에서 두 번 이상 나타나느 각 줄의 카운트와 텍스트를 출력한다.
// 이 프로그램은 표준 입력이나 파일 목록에서 읽는다.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			defer f.Close()
			if countLines(f, counts) {
				fmt.Println("dup filename:", arg)
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) bool {
	input := bufio.NewScanner(f)
	isdupFile := false
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			isdupFile = true
		}
	}
	// NOTE: input.Err()에서의 잠재적 오류는 무시한다

	return isdupFile
}
