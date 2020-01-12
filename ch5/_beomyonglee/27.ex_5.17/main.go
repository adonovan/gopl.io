package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func ElementsByTag(n *html.Node, tags ...string) []*html.Node {
	nodes := make([]*html.Node, 0)
	keep := make(map[string]bool, len(tags))
	for _, t := range tags {
		keep[t] = true
	}

	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		_, ok := keep[n.Data]
		if ok {
			nodes = append(nodes, n)
		}
		return true
	}
	forEachElement(n, pre, nil)
	return nodes
}

func forEachElement(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	u := make([]*html.Node, 0)
	u = append(u, n)
	for len(u) > 0 {
		n = u[0]
		u = u[1:]
		if pre != nil {
			if !pre(n) {
				return n
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
		if post != nil {
			if !post(n) {
				return n
			}
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: ex5.17 HTML_FILE TAG ...")
	}
	filename := os.Args[1]
	tags := os.Args[2:]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	for _, n := range ElementsByTag(doc, tags...) {
		fmt.Printf("%+v\n", n)
	}
}

/*
go build main.go
./main ongoing.html div
*/
