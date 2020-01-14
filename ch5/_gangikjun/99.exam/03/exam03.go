// findlinks1은 표준 입력에서 읽어 들인 HTML 문서 안의 링크를 출력한다
package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		log.Fatalln("http.Get : ", err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("html.Parse: %v\n", err)
	}

	printTextElement(doc)
}

func printTextElement(n *html.Node) {
	if n.Type == html.TextNode {
		if n.Data != "script" && n.Data != "style" {
			fmt.Print(n.Data)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printTextElement(c)
	}
}
