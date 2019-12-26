// basename은 디렉토리 구성 요소와 .suffix를 제거한다.
// 예) a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
	//NOTE: ignoring potential errors form input.Err()
}

func basename(s string) string {
	// 마지막 '/'와 그 전의 문자를 모두 제거한다.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	//마지막 '.' 전의 모든 문자를 남긴다.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
