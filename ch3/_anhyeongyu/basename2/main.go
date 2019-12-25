package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // "/"가 없으면 -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Println(basename("a/b/c.go"))
}
