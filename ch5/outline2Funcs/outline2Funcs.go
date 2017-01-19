package outline2Funcs

import (
	"fmt"
	"golang.org/x/net/html"
)

func PrintNodeType(node *html.Node) {

	switch {
	case node.Type == 0:
		fmt.Printf("Nodetype=ErrorNode\n")
	case node.Type == 1:
		fmt.Printf("Nodetype=TextNode\n")
	case node.Type == 2:
		fmt.Printf("Nodetype=DocumentNode\n")
	case node.Type == 3:
		fmt.Printf("Nodetype=ElementNode\n")
	case node.Type == 4:
		fmt.Printf("Nodetype=CommentNode\n")
	case node.Type == 5:
		fmt.Printf("Nodetype=DoctypeNode\n")
	}

	return

}

// Print URL for Links
//if n.Type == 3 && n.Data == "a" {
//	for _, a := range n.Attr {
//		if a.Key == "href" {
//			fmt.Printf("href        :%*s<%s>\n", depth*2, "", a.Val)
//		}
//	}
//}

func PrintRelatives(node *html.Node) {
	fmt.Printf("Parent: %v\n", node.Parent)
	fmt.Printf("FirstChild: %v\n", node.FirstChild)
	fmt.Printf("LastChild: %v\n", node.LastChild)
	fmt.Printf("PrevSibling: %v\n", node.PrevSibling)
	fmt.Printf("NextSibling: %v\n", node.NextSibling)

}
