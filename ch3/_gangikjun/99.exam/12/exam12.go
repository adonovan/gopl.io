package main

import (
	"fmt"
)

// isAngrams 아나그램이라면 true 반환
func isAngrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	r1 := make([]rune, 0, len(s1))
	r2 := make([]rune, 0, len(s2))

	for _, v := range s1 {
		r1 = append(r1, v)
	}
	for _, v := range s2 {
		r2 = append(r2, v)
	}

	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[len(r2)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAngrams("1234", "4321"))
	fmt.Println(isAngrams("1234", "3321"))

	fmt.Println(isAngrams("ㄱㄴㄷㄹㅁ", "ㅁㄹㄷㄴㄱ"))
	fmt.Println(isAngrams("ㄱㄴㄷㄹㅁ", "ㅁsㄷㄴㄱ"))
}
