package main

import (
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

// 크롤링 작업,
// 이 예제에서는 main 에서 crawl을 고루틴으로 마구 호출하면서
// 과도한 네트워크 연결을 동시에 수행하기 때문에, 작업이 실패할 수 있다

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)

	go func() { worklist <- []string{"http://gopl.io"} }()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
