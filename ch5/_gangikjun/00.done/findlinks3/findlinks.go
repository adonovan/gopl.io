package main

import (
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

func main() {
	breadthFirst(crawl, []string{"https://golang.org"})
}

// breadthFirst 는 작업 목록의 각 항목으로 f를 호출한다
// f에서 반환된 항목은 worklisk에 추가된다
// f는 한 항목에 최소한 한 번 실행된다
func breadthFirst(f func(item string) []string, worklisk []string) {
	seen := make(map[string]bool)
	for len(worklisk) > 0 {
		items := worklisk
		worklisk = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklisk = append(worklisk, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
