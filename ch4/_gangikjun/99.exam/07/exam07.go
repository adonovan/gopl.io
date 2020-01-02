package main

import (
	"fmt"
	"unicode/utf8"
)

// 할당안하고 할 수 있는건가요..?

func main() {
	b := []byte("안 녕 하 세 요")
	fmt.Println(string(reverseBytes(b)))
}

func reverseBytes(a []byte) []byte {
	if utf8.RuneCount(a) == 1 {
		return a
	}
	_, s := utf8.DecodeRune(a)
	return append(reverseBytes(a[s:]), a[:s]...)
}

// func reverse(b []byte) []byte {
// 	revSize := 0
// 	result := make([]byte, 0, 0)
// 	for i, v := range string(b) {
// 		revSize += utf8.RuneLen(v)

// 		for j := 0; j < utf8.RuneLen(v); j++ {
// 			//b[i], b[len(b)-revSize+j] = b[len(b)-revSize+j], b[i]
// 			result = append(result, b[i+j])
// 		}
// 		fmt.Println(string(result))
// 	}

// 	return result
// }
