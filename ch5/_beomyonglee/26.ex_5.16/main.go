package main

import (
	"bytes"
	"fmt"
)

func join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	b := bytes.Buffer{}
	for _, s := range strs[:len(strs)-1] {
		b.WriteString(s)
		b.WriteString(sep)
	}
	b.WriteString(strs[len(strs)-1])
	return b.String()
}

func main() {
	fmt.Println(join("/", "hi", "there"))
	fmt.Println(join("/"))
}
