// Fetchall은 url을 병렬로 반입하고 시간과 크기를 보고한다
package main

import (
	"bufio"
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
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	f, err := os.Create("result.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	for range os.Args[1:] {
		w := bufio.NewWriter(f)
		if _, err := w.WriteString(<-ch + "\n"); err != nil {
			log.Fatalln(err)
		}
		w.Flush()
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // ch 채널로 송신
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 리소스 누출 방지
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
