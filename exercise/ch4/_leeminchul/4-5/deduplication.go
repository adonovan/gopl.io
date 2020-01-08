package main

import "fmt"

func deduplicate(strings []string) []string {
	length := len(strings)
	if length == 0 { // nil slice 일 수 있음
		return strings
	}

	var result []string
	before := strings[0]
	for i, s := range strings[1:] {
		if before == s {
			result = strings[:i+1]
			break
		}
		before = s
	}
	if result == nil { // 중복이 없음
		return strings
	}

	for i := len(result) + 1; i < length; i++ {
		s := strings[i]
		if before != s {
			result = append(result, s)
			before = s
		}
	}
	return result
}

func main() {
	s := []string{"a", "a", "b", "c", "c", "d", "e", "f", "g", "g"}
	fmt.Printf("%p: %#v\n", s, s)
	s = deduplicate(s)
	fmt.Printf("%p: %#v\n", s, s) // []string{"a", "b", "c", "d", "e", "f", "g"}
}
