package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("1230000.1232"))
}

func comma(s string) string {
	var buf bytes.Buffer

	parts := strings.Split(s, ".")

	pos := 0

	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		fmt.Println("ACK " + parts[0][:pos])
		buf.WriteString(parts[0][:pos])
		buf.Write([]byte{','})
	}

	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write([]byte{','})
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}
	return buf.String()
}
