// echo2

package main

import (
	"fmt"
	"os"
)

func main() {
	// s, sep := "", ""
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

}

/*
go build -o echo2.exe
echo2 What a wonderful world


*/
