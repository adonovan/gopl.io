package main

import (
	"bytes"
	"fmt"
)

func main() {
	// comma("1234")
	fmt.Println(comma("12345678900"))
}

func comma(s string) string {
	var buf bytes.Buffer

	// fmt.Println(s)

	n := len(s)
	if n <= 3 {
		return s
	}

	r := n % 3
	buf.WriteString(s[:r])
	s = s[r:]

	for i, v := range s {

		if i%3 == 0 {
			buf.WriteString(",")
		}

		fmt.Fprintf(&buf, "%c", v)

	}

	return buf.String()
}
