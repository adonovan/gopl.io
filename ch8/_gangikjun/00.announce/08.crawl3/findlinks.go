package main

import (
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

// crawl2 번 예제와는 다른 방법으로 동시성 문제 해결
// 세마포어 없이 하는 대신, 20개의 고루틴을 호출한다

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // URL 목록. 중복이 있을 수 있음
	unseenLinks := make(chan string) // URL 중복 제거

	go func() { worklist <- []string{"http://gopl.io"} }()

	// 20개의 고루틴을 생성해 찾아보지 않은 항목을 탐색
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// 메인 고루틴은 작업 목록의 항목에서 중복을 제거하고
	// 아직 찾아보지 않은 항목을 크롤러에 보낸다
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
