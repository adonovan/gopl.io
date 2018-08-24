// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// ex 1.1 solution
// print also Args[0] the name of the command invoked

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("command:", os.Args[0])
	fmt.Println("args:", strings.Join(os.Args[1:], " "))
}
