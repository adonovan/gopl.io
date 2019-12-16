package main

import "fmt"

func main() {
	fmt.Println(comma("1234"))
	fmt.Println(comma("100000000000"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
