package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

// breadthFirst는 작업 목록의 각 항목으로 f를 호출한다.
// f에서 반환된 항목은 worklist에 추가된다.
// f는 한 항목에 최소한 한 번 실행한다.
func breadthFirst(f func(item string) []string, worklist []string) {
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
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}

/*
go build main.go
./main https://golang.org
*/
