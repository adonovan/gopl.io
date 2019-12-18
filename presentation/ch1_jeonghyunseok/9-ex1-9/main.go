// 1.9 fetch를 수정해 resp.Status에 있는 HTTP 응답 코드도 같이 출력하라.
// start with fatch.go

// fetch is
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("Status: %s. %s\n", resp.Status, b)
	}
}

/*
go build -o ex1-9.exe
ex1-9.exe http://www.president.go.kr/
ex1-9.exe http://www.bad.gopl.io

*/
