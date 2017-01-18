// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	//"github.com/dugwill/gopl.io/ch5/outline2Funcs"
	"golang.org/x/net/html"
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

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

func startElement(n *html.Node) {
	//if n.Type == html.ElementNode {

	if n.Type == html.TextNode && (n.Data == "" || n.Data == "\n  " || n.Data == "\n" || n.Data == "\n\n") {
		return
	}

	fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)

	//outline2Funcs.PrintNodeType(n)

	//if n.Type == html.TextNode {
	//	fmt.Printf("ASCII: ")
	//	for i := 0; i < len(n.Data); i++ {
	//		fmt.Printf("%x ", n.Data[i])
	//	}
	//}
	depth++
	//fmt.Println()

}

func endElement(n *html.Node) {

	if n.Type == html.TextNode && (n.Data == "" || n.Data == "\n  " || n.Data == "\n" || n.Data == "\n\n") {
		return
	}

	depth--

	if n.Type != html.TextNode {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}

}

//!-startend
