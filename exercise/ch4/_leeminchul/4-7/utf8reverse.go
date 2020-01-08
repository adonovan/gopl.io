package main

import (
	"fmt"
)

func Utf8Reverse(bytes []byte) {
	reverse(bytes)

	var size int
	for i := len(bytes); 0 < i; i -= size {
		size = utf8Len(bytes[i-1])
		reverse(bytes[i-size : i])
	}
}

func utf8Len(firstByte byte) int {
	firstByte >>= 3
	if firstByte == 0b11110 {
		return 4
	}
	firstByte >>= 1
	if firstByte == 0b1110 {
		return 3
	}
	firstByte >>= 1
	if firstByte == 0b110 {
		return 2
	}
	return 1
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []byte("가나다라마바사아자차카타파하")
	Utf8Reverse(s)
	fmt.Println(string(s))
}
