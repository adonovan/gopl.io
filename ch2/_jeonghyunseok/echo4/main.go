package main

import (
	"flag"
	"fmt"
	"strings"
)

// 플래그, 초기값, 설명
var n = flag.Bool("n", false, "omit trailing newline") // true 이면 줄바꿈
var sep = flag.String("s", " ", "seperator")           // seperator 설정

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

/*
go build -o echo4.exe
echo4.exe a bc def ghij
echo4.exe -n=true a bc def ghij
echo4.exe -s="$" a bc def ghij
*/
