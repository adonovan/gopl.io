// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

//Modified by Douglas Will
//excercise 5.12 the instructions were modify the 'outline' example and
//turn the 'startElement' and 'endElement' functions
//into anonymous functions based within the 'outline' function.
//the 'startElement' and 'endElement' functions are now variables of type 'func(n *html.node)''
//Those variables are assigned to anonymous funcs with the steps of the original functions.
//The are then passed, as variables, to the 'forEachNode' function
//Both Anonymous functions share access to the var 'depth' which is now local to the 'outline' function

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

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

	var depth int // now local to the outline function

	// declare the vars to pass to func 'foreachnode'
	var startElement, endElement func(n *html.Node)

	//assign var 'startElement' to an anonymous function
	startElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}

	//assign var 'endElelement' to an anonymous function
	endElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)

		}
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
