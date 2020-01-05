// Findlinks3 은 웹 크롤링을 한다. 명령줄의 URLs 로 시작해서

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gopl.io/ch5/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	// item 이라는 문자열을 받아서 문자열 슬라이스를 리턴하는 함수와, worklist 라는 문자열 슬라이스

	seen := make(map[string]bool)
	for len(worklist) > 0 {
		// worklist 를 받아서 items 로 옮기고
		// worklist 는 nil
		items := worklist
		worklist = nil

		for _, item := range items {

			// 그리고 해당 item 이 seen 이 아닌 경우만
			// item 을 함수에 넣고 그 결과값을 worklist 에 추가해줌
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
links.Extract() 는 url 을 받아서 해당 url 에 있는 url 들을 모두 문자열로 리턴해주는 것으로 보인다.
결국 crawl 이 하는 일이 그거다. 

breadthFirst() 는 이 crawl 함수를 받아서 seen 이 false 인 경우만 계속 worklist 에 추가하여 
다시금 crawl() 을 돌린다. 

*/

/*
go build -o findlinks3.exe
findlinks3.exe http://www.naver.com

*/
