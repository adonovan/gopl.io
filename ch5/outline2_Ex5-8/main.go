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
	//"github.com/dugwill/gopl.io/ch5/outline2Funcs"
)

var depth int
var cont = true

func main() {

	outline(os.Args[1], os.Args[2])
}

func outline(url, id string) error {
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
	if cont {
		cont := forEachNode(doc, id, startElement, endElement)
	}
	//!-call

	return nil
}

//func ElementByID() {}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, id string, pre func(n *html.Node, id string) bool, post func(n *html.Node)) bool {
	if pre != nil {
		pre(n, id)
	}

	//fmt.Printf("NodeType: %v\n",n.Type)
	//printNodeType(n)

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		if cont {
			cont := forEachNode(c, id, startElement, endElement)
		}

		if post != nil {
			post(n)
		}
	}
	return cont
}

//!-forEachNode

//!+startend
//var depth int

func startElement(n *html.Node, id string) bool {

	if n.Type == html.TextNode && (n.Data == "" || n.Data == "\n  " || n.Data == "\n" || n.Data == "\n\n") {
		return true
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
			for _, a := range n.Attr {
				fmt.Printf(" %s=", a.Key)
				fmt.Printf("\"%s\"", a.Val)
				if id == a.Key || id == a.Val {
					cont = false
				}
			}

		}

		//if n.FirstChild != nil && (n.FirstChild.Type == html.TextNode && (n.FirstChild.Data == "" || n.FirstChild.Data == "\n  " || n.FirstChild.Data == "\n" || n.FirstChild.Data == "\n\n")) {
		//	fmt.Printf(">\n")
		//} else {
		//	fmt.Printf(">\n")
		//}

		fmt.Printf(">\n")

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

	return cont

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
