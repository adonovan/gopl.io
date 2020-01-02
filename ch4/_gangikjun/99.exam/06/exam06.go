package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	b := []byte("안 녕 하 세 요")
	fmt.Println(removeSpace(b))
	fmt.Println(string(removeSpace(b)))
}

func removeSpace(b []byte) []byte {
	var buf bytes.Buffer
	s := string(b)

	for _, r := range s {
		if unicode.IsSpace(r) {
			continue
		}
		buf.WriteRune(r)
	}

	return buf.Bytes()
}
