// findlinks1은 표준 입력에서 읽어 들인 HTML 문서 안의 링크를 출력한다
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	htmlText, err := getHTML("https://golang.org")
	if err != nil {
		log.Fatalln("getHTML : ", err)
	}

	htmlReader := bytes.NewReader(htmlText)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		log.Fatalf("findlnks1: %v\n", err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func getHTML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}
