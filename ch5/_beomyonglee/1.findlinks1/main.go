// Findlinks1은 표준 입력에서 읽어 들인 HTML 문서 안의 링크를 출력한다.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Goland에서는 아래에서 Parse 함수와 같이 몇몇 빨간색으로 표시되어 있지만
// 실행에는 문제가 없음
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit은 n에서 찾은 각 링크를 links에 추가하고 결과를 반환한다.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

/*
이 파일이 있는 위치에서 다음 명령어를 실행한다.

$ go build gopl.io/ch1/fetch
$ go build main.go

./fetch https://golang.org | ./main
/
/doc/
/pkg/
/project/
/help/
/blog/
https://play.golang.org/
/dl/
https://tour.golang.org/
https://blog.golang.org/
/doc/copyright.html
/doc/tos.html
http://www.google.com/intl/en/policies/privacy/
http://golang.org/issues/new?title=x/website:
https://google.com
*/
