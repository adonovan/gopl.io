// Package links 는 링크 추출 함수를 제공한다
package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Extract 는 지정된 URL로 HTTP GET 요청을 생성하고
// 결과를 HTML응답으로 파싱한 후 HTML문서 내의 링크를 반환한다
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
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
