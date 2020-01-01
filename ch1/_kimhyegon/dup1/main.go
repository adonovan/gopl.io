package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	i := 0
	for input.Scan() {
		// debug
		fmt.Printf("input value: [%d],%v\n", i, input.Text())
		counts[input.Text()]++
	}

	// NOTE: input.Err() 에서의 잠재적 오류는 무시한다.
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
