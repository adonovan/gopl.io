package main

import "fmt"

func main() {
	// s := "abc"
	// b := []byte(s)
	// s2 := string(b)

	// fmt.Println(s2)
	fmt.Println(comma("01234567890000"))
}

func comma(s string) string {
	fmt.Println(s)
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
