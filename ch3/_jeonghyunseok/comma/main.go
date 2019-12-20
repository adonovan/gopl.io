// Comma 아규먼트를 3자리마다 콤마를 넣어 출력한다.
// 재귀 함수사용.

package main

import (
	"fmt"
	"os"
)

func main() {
	for _, a := range os.Args[1:] {
		fmt.Printf("  %s\n", comma(a))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

/*
go build -o comma.exe
comma.exe 123 1234 123456 12345678

*/

// 123
// 1,234
// 123,456
// 12,345,678
