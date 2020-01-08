package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s // strings 이 변질됨
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // 길이가 0인 원본 slice
	for _, s := range strings {
		if s != "" {
			out = append(out, s) // strings 이 변질됨
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // ["one" "three"]
	fmt.Printf("%q\n", data)           // ["one" "three" "three"]

	data = []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data)) // ["one" "three"]
	fmt.Printf("%q\n", data)            // ["one" "three" "three"]
}
