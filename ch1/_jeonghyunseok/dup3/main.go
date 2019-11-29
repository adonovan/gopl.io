// Dup3 prints count of line appear
// go build -o dup3.exe
// dup3 black5line.txt
// dup3 tmp.txt -> strange good 2
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			// fmt.Println(line)
			counts[line]++
		}
	}

	for lineText, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, lineText)
			fmt.Printf("%d\t%x\n", n, []byte(lineText))
		}
	}
}
