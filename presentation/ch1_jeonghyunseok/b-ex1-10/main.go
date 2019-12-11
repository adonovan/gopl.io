// 1.10 많은 양의 데이터를 생성하는 웹사이트를 찾아라.
// fetchall을 두 번 연속으로 실행하고 결과 시간이 얼마나 달라지는지를 통해 캐시 여부를 조사하라.
// 매번 같은 내용을 받는가?
// 이 결과를 조사하기 위해 fetchall이 결과를 파일로 출력하게 수정하라.
// start with fetchall.go

// fetchall
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	f, err := os.Create("result.txt")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	for _, url := range os.Args[1:] {
		go fetch(url, ch, f)
		go fetch(url, ch, f)
	}
	for range os.Args[1:] {
		fmt.Print(<-ch) // blocking
		fmt.Print(<-ch) // blocking
	}
	result := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	f.WriteString(result)
}

func fetch(url string, ch chan<- string, file *os.File) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		result := fmt.Sprint(err)
		file.WriteString(result)
		ch <- result
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		result := fmt.Sprintf("while reading %s: %v\n", url, err)
		file.WriteString(result)
		ch <- result
		return
	}
	secs := time.Since(start).Seconds()
	result := fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
	file.WriteString(result)
	ch <- result
}

/*
go build -o ex1-10.exe
ex1-10.exe https://golang.org http://www.naver.com http://www.bmw.co.jp/ja/index.html
type result.txt

no cache I guess
*/
