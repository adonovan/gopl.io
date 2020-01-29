// Crawl1 은 웹 링크를 크롤링한다. cli 아규먼트부터 시작해서 말이다
// fd 를 빨리 소진한다.
// worklist 를 닫지 않기에 끝나지도 않는다.package crawl1

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gopl.io/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// 하나 이상의 url 을 cli 에서 받아서 worklist 에 넣어준다.
	worklist := make(chan []string)
	go func() { worklist <- os.Args[1:] }()  

	seen := make(map[string]bool)

	// worklist 채널을 닫아주는게 없으니 for 문은 무한루프가 될 것이다. 
	for list := range worklist {
		for _, link := range list {  // 여기서는 나는 http://naver.com 하나만 넣어준 것임
			if !seen[link] {  // 한번 방문한 link 는 방문 안하도록 해둠
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)  // crawl 한 결과, 즉, 링크에 있는 링크들을 worklist 에 다시 넣어줌 
				}(link)
			}
		}
	}
}

/*
go build -o fl.exe
fl http://www.naver.com

// 이게 아마도 file description 이 바닥난 것이겠지?
panic: too many concurrent operations on a single file or socket (max 1048575)
goroutine 1119058 [running]:

*/
