package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

InputScan:
	for {
		if !input.Scan() {
			if input.Err() != io.EOF {
				fmt.Printf("input.Scan(): %v\n", input.Err())
			}
			break InputScan
		}
		fmt.Println(basename(input.Text()))
	}
	fmt.Println("program terminated")
}

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

/*

go build -o basename1.exe
basename1.exe
a
a.go
a/b/c.go

*/
