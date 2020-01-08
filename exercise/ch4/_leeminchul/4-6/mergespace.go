package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func mergeSpace(bytes []byte) []byte {
	length := len(bytes)
	result := bytes[:0]
	var size int
	var char rune
	var spaceIndex, beforeSpaceIndex int
	beforeWordIndex := 1
	isSpace := false
	beforeIsSpace := false

	for i := 0; i < length; i += size {
		isSpace = false
		char, size = utf8.DecodeRune(bytes[i:]) // 첫 글자를 decode
		if size == 1 || size == 2 {
			isSpace = unicode.IsSpace(char)
		}

		if isSpace && !beforeIsSpace { // 공백이 시작 될 때
			spaceIndex = i
			beforeIsSpace = true
		} else if !isSpace && beforeIsSpace { // 공백이 끝날 때
			bytes[i-1] = ' '
			result = append(result, bytes[beforeWordIndex-1:spaceIndex]...)
			beforeSpaceIndex = spaceIndex
			beforeWordIndex = i
			beforeIsSpace = false
		}
	}

	if beforeIsSpace {
		result = append(result, bytes[beforeWordIndex-1:spaceIndex+1]...)
	} else {
		result = append(result, bytes[beforeSpaceIndex:]...)
	}
	return result
}

func main() {
	b := []byte("    독 \t도\xc2\xa0는 우리 \xc2\x85땅  ")
	b = mergeSpace(b)

	fmt.Printf("[%s]\n", string(b))
}
