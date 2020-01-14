// findlinks1은 표준 입력에서 읽어 들인 HTML 문서 안의 링크를 출력한다
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
		log.Fatalf("html.Parse: %v\n", err)
	}

	element := make(map[string]int)
	setElement(element, doc)
	for k, v := range element {
		fmt.Println(k + ":" + strconv.Itoa(v))
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

func setElement(element map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		element[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		setElement(element, c)
	}
}
