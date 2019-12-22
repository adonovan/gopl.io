// Noempty 는 인라인 슬라이스 알고리즘의 예제이다.
package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// strings 는 유명한 패키지 이름이지만 그냥 변수로 사용
func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	fmt.Println("--")

	data2 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data2))
	fmt.Printf("%q\n", data2)
}

// go run main.go

// ["one" "three"]
// ["one" "three" "three"]
// --
// ["one" "three"]
// ["one" "three" "three"]
