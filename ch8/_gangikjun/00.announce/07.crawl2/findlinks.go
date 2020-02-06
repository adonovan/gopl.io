package main

import (
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

// 한번에 최대 20번의 호출만 일어나도록 수정

// tokens는 최대 20개의 동시 요청을 강제하는
// 카운팅 세마포어
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 토큰 획득
	list, err := links.Extract(url)
	<-tokens // 토큰 해제
	if err != nil {
		log.Print(err)
	}
	return list
}

// 아래의 main에서는 초기 URL에서 도달 가능한 모든 링크를 찾은 후에도 종료되지 않고 있다
// func main() {
// 	worklist := make(chan []string)

// 	go func() { worklist <- []string{"http://gopl.io"} }()

// 	seen := make(map[string]bool)
// 	for list := range worklist {
// 		for _, link := range list {
// 			if !seen[link] {
// 				seen[link] = true
// 				go func(link string) {
// 					worklist <- crawl(link)
// 				}(link)
// 			}
// 		}
// 	}
// }

// 작업 목록이 비어 있고 활성화된 크롤링 고루틴이 없을 때 메인 루프에서 나오도록 수정
func main() {
	worklist := make(chan []string)
	var n int // 작업 목록에 대기 중인 작업 개수

	n++
	go func() { worklist <- []string{"http://gopl.io"} }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
