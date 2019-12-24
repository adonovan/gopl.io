package main

import (
	"bytes"
	"fmt"
)

// intToString 은 fmt.Sprint(values)와 유사하지만 콤마를 추가한다
func intToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteRune(']')
	return buf.String()
}

func main() {
	fmt.Println(intToString([]int{1, 2, 3}))
}
