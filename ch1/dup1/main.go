// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	//fmt.Println(input)
	//var value int = 0
	//value := 0
	for input.Scan() {
		if input.Text() == "end" { break }
		//fmt.Println("input", input.Text())
		counts[input.Text()]++
		//fmt.Println(counts)
	}
	// NOTE:added potential error check
	if input.Err() != nil{
        fmt.Print("err")
    }
	fmt.Println(counts)
	for k, c := range counts {
		if c>1 {
			fmt.Printf("%s\t%d\n",k,c)
		}
	}
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }
}

//!-
