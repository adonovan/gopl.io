package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func main() {

}

// soleTitle 은 비어있지 않은 첫 번째 title 원소의 텍스트를 반환하며,
// 한 개가 아니라면 error를 반환한다
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// 패닉이 아님
		case bailout{}:
			// 예상된 패닉
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // 예상치 못한 패닉; 패닉 지속
		}
	}()

	// 비어있지 않은 title이 두 개 이상이면 재귀에서 탈출한다.
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // 여러 title 원소
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

// forEachNode 는 루트 노드 n에 있는 각 노드 x에 대해
// pre(x) 함수와 post(x) 함수를 호출한다. 두 함수는 옵션이다.
// pre는 하위 노드를 방문하기 전에 호출되며
// post는 하위 노드를 방문한 후에 호출된다
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
