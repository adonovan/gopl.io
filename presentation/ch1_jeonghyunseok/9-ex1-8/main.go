// 1.8 각 인수 URL에 http:// 접두사가 누락된 경우 이를 추가하도록 fetch를 수정하라.
// strings.HasPrefix를 사용할 수도 있다. https://golang.org/pkg/strings/#HasPrefix
// start with fatch.go

// fetch is
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	// ignore
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			fmt.Println("no prefix!", url)
			continue
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n----------\n", err)
			continue
			// os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("url: %s\n%s\n----------\n", url, b)
	}
	fmt.Println()
	// add prefix
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("url: %s\n%s\n----------\n", url, b)
	}
}

/*
go build -o ex1-8.exe
ex1-8.exe http://www.president.go.kr/ www.google.com


*/
