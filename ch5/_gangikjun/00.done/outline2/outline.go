package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

var depth int

func main() {
	outline("http://gopl.io")
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
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

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
