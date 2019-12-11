package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// ignore input.Err()
	fmt.Println(counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/*
go build -o dup1.exe
dup1.exe

<type and enter>
// Ctrl+Z and Enter will stop "for loop"
*/

// aa
// aa
// bb
// bb
// ^Z
// map[aa:2 bb:2]
// 2       aa
// 2       bb
