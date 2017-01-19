package main

	import (
	"fmt"
	"os"
	//"net/http"
	"golang.org/x/net/html"
	"strings"
	"github.com/dugwill/myProjects/myHTML"
	)


	
func main(){

var numWords, numImages int	

	var err error

		for _, url := range os.Args[1:] {

		doc:=myHTML.GetHTML(url)
		numWords, numImages = countWordsAndImages(doc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "HTMLCount: %v\n", err)
			continue
		}

	}	

	fmt.Printf("Number of words: %d\n",numWords)
	fmt.Printf("Number of images: %d\n",numImages)

} //End func main

func countWordsAndImages(n *html.Node) (words, images int){
	if n.Type == html.TextNode {
			//fmt.Printf("n.Type: %v \n",n.Type)
			words += len(strings.Split(n.Data, " "))
	}

	if n.Data == "img" {
		//fmt.Printf("n.Type: %v \tn.Data: %v\n",n.Type, n.Data)
		images++
	} else {
		if n.FirstChild != nil {
			w, i := countWordsAndImages(n.FirstChild)
			words += w
			images += i
		}
		if n.NextSibling != nil {
			w, i := countWordsAndImages(n.NextSibling)
			words += w
			images += i
		}
	}
	return
}
