package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	seen := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	text := ""
	for input.Scan() {
		text = input.Text()
		fmt.Println(text)
		seen[text]++
	}
	if err := input.Err(); err != nil {
		log.Fatalln("input :", err)
	}

	for k, v := range seen {
		fmt.Printf("word : %s\t, count : %d\n", k, v)
	}
}
