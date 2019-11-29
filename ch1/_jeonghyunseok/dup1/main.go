package main

import (
	"bufio"
	"fmt"
	"os"
)

// dup1
// Ctrl+Z and Enter will stop "for loop"
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// ignore input.Err()
	fmt.Println(counts)
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
