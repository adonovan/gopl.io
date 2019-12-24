package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var temp string
	var decimalPlace string

	// under
	dot := strings.LastIndex(s, ".")
	if dot >= 0 {
		temp = s[:dot]
		decimalPlace = s[dot:]
	} else {
		temp = s
	}

	n := len(temp)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i, v := range temp {
		if i != 0 && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(v)
	}

	return buf.String() + decimalPlace
}

func main() {
	fmt.Println(comma("1123456789.0262345"))
}
