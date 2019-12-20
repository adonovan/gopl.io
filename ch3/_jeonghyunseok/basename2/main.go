// Basename2 는 파일이름을 표준 입력으로 읽고, 베이스네임을 출력한다.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	// if dot := strings.LastIndex(s, "."); dot >=0 {
	// 	s = s[:dot]
	// }

	dot := strings.LastIndex(s, ".")
	fmt.Printf("dot type %T, value: %v\t", dot, dot)

	if dot >= 0 {
		s = s[:dot]
	}
	return s
}

/*
go build -o basename2.exe
basename2.exe
a
a.go
a.b.c.go
a/b/c/d/e.go

*/
