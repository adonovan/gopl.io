package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin) // Parse 를 해주면 *Node 라는 놈을 리턴한다.
	if err != nil {
		log.Fatalf("fintlinks1: %v\n", err)
	}
	fmt.Println(*doc)
	fmt.Println("--")
	for _, link := range visit(nil, doc) { // *Node 를 visit 에 넣어주면 []string 을 회신해준다.
		log.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" { // Node 의 타입이 Element 이고, a 인 경우만
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
go build -o findlink1.exe
findlink1.exe < test.html

*/

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
