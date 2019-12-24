package main

import (
	"fmt"
	"os"
)
import "../links"

func breathFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		fmt.Print(err)
	}
	return list
}
func main() {
	breathFirst(crawl, os.Args[1:])
}