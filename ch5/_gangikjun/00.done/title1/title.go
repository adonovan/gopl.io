package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	title("http://gopl.io")
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Contetn-Type이 HTML인지 검사한다(ex, "text/html; charset=utf-8")
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}

	forEachNode(doc, visitNode, nil)
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
