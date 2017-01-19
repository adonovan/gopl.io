// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"

	"github.com/dugwill/gopl.io/ch5/outline2Funcs"
)

var depth int

func main() {

	for _, url := range os.Args[1:] {
		outline(url)
	}
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

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	//fmt.Printf("NodeType: %v\n",n.Type)
	//printNodeType(n)

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
//var depth int

func startElement(n *html.Node) {

	if n.Type == html.TextNode && (n.Data == "" || n.Data == "\n  " || n.Data == "\n" || n.Data == "\n\n") {
		return
	}

	//outline2Funcs.PrintNodeType(n)

	if n.Type == html.ElementNode {

		switch {
		case n.FirstChild == nil:
			fmt.Printf("%*s<%s/", depth*2, "", n.Data)
		case n.FirstChild != nil:
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
		}
		//outline2Funcs.PrintNodeType(n)
		if n.Attr != nil {
			outline2Funcs.PrintAttributes(n)
		}

		if n.FirstChild != nil && (n.FirstChild.Type == html.TextNode && (n.FirstChild.Data == "" || n.FirstChild.Data == "\n  " || n.FirstChild.Data == "\n" || n.FirstChild.Data == "\n\n")) {
			fmt.Printf(">\n")
		} else {
			fmt.Printf(">\n")
		}

		//fmt.Printf(">\n")

	}

	if n.Type == html.TextNode {

		if n.Data != "" && n.Data != "\n  " && n.Data != "\n" {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)

			//	for i := 0; i < len(n.Data); i++ {
			//		fmt.Printf("%x ", n.Data[i])
			//	}

		}
	}

	//if n.Type != html.ElementNode {
	//	fmt.Printf("\n")
	//}

	depth++

}

func endElement(n *html.Node) {

	if n.Type == html.TextNode && (n.Data == "" || n.Data == "\n  " || n.Data == "\n" || n.Data == "\n\n") {
		return
	}

	depth--

	if n.Type == html.ElementNode {

		switch {
		case n.FirstChild == nil:
			fmt.Printf("")
		case n.FirstChild != nil:
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}

	}

}

//!-startend
