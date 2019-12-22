// Printints 는 bytes.Buffer 를 사용하는 법을 보여준다.

package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteString("]")
	return buf.String()
	// buf.WriteByte('[')
	// for i, v := range values {
	// 	if i > 0 {
	// 		buf.WriteString(", ")
	// 	}
	// 	fmt.Fprintf(&buf, "%d", v)
	// }
	// buf.WriteByte(']')
	// return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}

/*
go run main.go
*/
