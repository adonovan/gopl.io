package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "abcdefghij"
	s2 := "jihgfedcba"

	fmt.Printf("아나그램 여부: %t\n", compareAnagrams(s1, s2))

}

func compareAnagrams(s1 string, s2 string) bool {

	if strings.Compare(s1, reverse(s2)) == 0 {
		return true
	}

	return false
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
