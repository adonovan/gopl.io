// Noempty는 슬라이스 직접 변경 알고리즘의 예다
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", noempty(data))
	fmt.Printf("%q\n", noempty2(data))
	fmt.Printf("%q\n", data)
}

func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func noempty2(strings []string) []string {
	out := strings[:0] // 길이가 0인 원본의 슬라이스
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
