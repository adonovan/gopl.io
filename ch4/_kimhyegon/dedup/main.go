package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sean := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !sean[line] {
			sean[line] = true
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup:%v", err)
		os.Exit(1)
	}
}
