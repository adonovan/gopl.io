package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	for i, v := range s {
		if i != 0 && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(v)
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("123456789"))
}
