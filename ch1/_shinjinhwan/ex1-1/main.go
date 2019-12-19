package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[0]) // printed running file name
	fmt.Println(strings.Join(os.Args[1:], " "))
}