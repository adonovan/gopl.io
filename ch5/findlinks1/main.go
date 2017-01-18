// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var layer=0

var nodeCount = make(map[string]int)


func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	PrtNodeCount()
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string { //visit take slice of strings and Ptr to HTML node

	if n.Type == html.ElementNode{
		NodeCount(n.Data)
	}

	/* The next statement looks for textnodes and prints the data*/
	//if n.Type == html.TextNode{
		//fmt.Printf("TextNode: %s\n",n.Data)
	//}


	if n.Type == html.ElementNode && n.Data == "img" { // if n is a html element node type A
			
		for _, a := range n.Attr { //iterate through all attributes

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

func NodeCount(nodeData string){

nodeCount[nodeData]++

}

func PrtNodeCount(){

	fmt.Printf("\nElementNode Counts\n")
	for ntype,count := range nodeCount{
		fmt.Printf("%s\t\t\t%d\n",ntype, count)
	}

}

//!-visit

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
