// Outline은 HTML 문서 개요 트리를 출력한다.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

/*
$ go build gopl.io/ch1/_BeomyongLee/fetch
$ go build main.go
$ ./fetch https://golang.org | ./main
[html]
[html head]
[html head meta]
[html head meta]
...
*/
