package main

import "fmt"

// comma 는 양의10진 정수 문자열에 콤마를 삽입한다
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println(comma("12345"))
}
