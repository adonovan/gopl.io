// crawl1 은 무한정 goroutine 이 생성되었지만
// 여기서는 buffered channel 을 이용하여 쎄마포어처럼 써서 열려있는 goroutine 개수를 조절했다.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gopl.io/ch5/links"
)

var tokens = make(chan struct{}, 20) // 20개의 버퍼를 가진 토큰

// 하나의 url 을 받아서 수많은 urls 를 리턴해준다.
func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}            // 여기가 열려있는 개수를 조절하는 부분. tokens 20개가 꽉차면 블록된다.
	list, err := links.Extract(url) // url 에 있는 links 를 추출해내서 list 에 담는다.
	<-tokens                        // 다시 tokens 하나를 풀어내준다.

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // 이건 무슨 개수를 셀까

	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- { // link 가 하나 추가될 때마다 n 이 늘어난다.
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

/*
go build -o fl.exe
fl.exe http://www.naver.com


*/
