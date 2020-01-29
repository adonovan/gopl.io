// Crawl3 은 bounded parallelism 을 사용한다.
// 종료 문제는 언급하지 말자.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gopl.io/ch5/links"
)

// crawl 은 특별할 것이 없다. 
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}


func main() {
	worklist := make(chan []string)  // 일단 이건 worklist 이지만 중복된 녀석이 마구 들어갈 수 있다.
	unseenLinks := make(chan string) // 이건 중복제거한 버전. 어떻게 이용될까/ 

	go func() { worklist <- os.Args[1:] }()

	// 크롤러는 딱 20개만 만들자. 
	// 20개의 goroutine 이 도는데 unseenLinks 채널로 값이 들어오면 이것들 중 하나가 받아서 처리할 것이다.
	// link 를 받으면 다시 worklist 로 넣어준다. 아마도 worklist 로 들어온걸 정리해서 unseenLinks 로 넣어주겠지?
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

    // link 가 unseen 인 경우만 채널로 넣어준다. 
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

/*
go build -o fl.exe
fl.exe http://www.lge.com


*/