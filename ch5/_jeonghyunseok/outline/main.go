// Outline 은 HTML 문서 트리의 아웃라인을 출력해준다

package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("outline: %v\n", err)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

/*
go build -o outline.exe
outline.exe < test.html


*/

// [html]
// [html head]
// [html body]
// [html body style]
// [html body a]
// [html body a img]
// [html body a]
// [html body a]
// [html body a]
// [html body a]
// [html body a]
// [html body h1]
// [html body div]
// [html body div img]
// [html body div div]
// [html body div]
// [html body p]
// [html body p span]
// [html body p br]
// [html body p span]
// [html body p span]
// [html body p br]
// [html body p br]
// [html body p br]
// [html body p span]
// [html body a]
// [html body a span]
// [html body a]
// [html body a span]
// [html body a]
// [html body a span]
