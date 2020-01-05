package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}

	reverse(a[:])

	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}

	reverse(s[:2])
	reverse(s[2:])
	reverse(s)

	fmt.Println(s)

	// Interactive test of reverse
	input := bufio.NewScanner(os.Stdin)

outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
}

func reverse(array []int) {
	for i, j := 0, len(array) -1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}