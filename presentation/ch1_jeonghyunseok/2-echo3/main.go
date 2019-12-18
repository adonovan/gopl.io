// Echo3
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

/*
go build -o echo3.exe
.\echo3 Who let the dogs out!
*/
