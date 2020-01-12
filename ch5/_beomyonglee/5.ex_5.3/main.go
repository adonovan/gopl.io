package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func printTagText(r io.Reader, w io.Writer) error {
	z := html.NewTokenizer(os.Stdin)
	var err error
	stack := make([]string, 20)
Tokenize:
	for {
		switch z.Next() {
		case html.ErrorToken:
			break Tokenize
		case html.StartTagToken:
			b, _ := z.TagName()
			stack = append(stack, string(b))
		case html.TextToken:
			cur := stack[len(stack)-1]
			if cur == "script" || cur == "style" {
				continue
			}
			text := z.Text()
			if len(strings.TrimSpace(string(text))) == 0 {
				continue
			}
			w.Write([]byte(fmt.Sprintf("<%s>", cur)))
			w.Write(text)
			if text[len(text)-1] != '\n' {
				io.WriteString(w, "\n")
			}
		case html.EndTagToken:
			stack = stack[:len(stack)-1]
		}
	}
	if err != io.EOF {
		return err
	}
	return nil
}

func main() {
	err := printTagText(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

/*
go build gopl.io/ch1/_BeomyongLee/fetch
go build main.go
./fetch https://golang.org | ./main
<title>The Go Programming Language
<a>Documents
<a>Packages
<a>The Project
...
*/
