package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int

type PrettyPrinter struct {
	w   io.Writer
	err error
}

func NewPrettyPrinter() PrettyPrinter {
	return PrettyPrinter{}
}

func (pp PrettyPrinter) Pretty(w io.Writer, n *html.Node) error {
	pp.w = w
	pp.err = nil
	pp.forEachNode(n, pp.start, pp.end)
	return pp.Err()
}

func (pp PrettyPrinter) Err() error {
	return pp.err
}

func (pp PrettyPrinter) forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	if pp.Err() != nil {
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		pp.forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
	if pp.Err() != nil {
		return
	}
}

func (pp PrettyPrinter) printf(format string, args ...interface{}) {
	_, err := fmt.Fprintf(pp.w, format, args...)
	pp.err = err
}

func (pp PrettyPrinter) startElement(n *html.Node) {
	end := ">"
	if n.FirstChild == nil {
		end = "/>"
	}

	attrs := make([]string, 0, len(n.Attr))
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}
	attrStr := ""
	if len(n.Attr) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	name := n.Data

	pp.printf("%*s<%s%s%s\n", depth*2, "", name, attrStr, end)
	depth++
}

func (pp PrettyPrinter) endElement(n *html.Node) {
	depth--
	if n.FirstChild == nil {
		return
	}
	pp.printf("%*s</%s>\n", depth*2, "", n.Data)
}

func (pp PrettyPrinter) startText(n *html.Node) {
	text := strings.TrimSpace(n.Data)
	if len(text) == 0 {
		return
	}
	pp.printf("%*s%s\n", depth*2, "", n.Data)
}

func (pp PrettyPrinter) startComment(n *html.Node) {
	pp.printf("<!--%s-->\n", n.Data)
}

func (pp PrettyPrinter) start(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		pp.startElement(n)
	case html.TextNode:
		pp.startText(n)
	case html.CommentNode:
		pp.startComment(n)
	}
}

func (pp PrettyPrinter) end(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		pp.endElement(n)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	pp := NewPrettyPrinter()
	pp.Pretty(os.Stdout, doc)
}

/*
go build gopl.io/ch1/_BeomyongLee/fetch
go build pretty.go
./fetch https://golang.org | ./gretty
<html lang="en">
  <head>
    <meta charset="utf-8"/>
    <meta name="description" content="Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."/>
...
*/
