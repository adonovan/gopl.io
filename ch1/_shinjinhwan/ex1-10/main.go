package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
	f, err := os.Create("resp-result.txt")
	if err != nil {
		panic(err)
	}

	for range os.Args[1:] {
		recv := <- ch
		fmt.Println(recv) // receive data through ch
		w, err := f.WriteString(recv + "\n")
		if err != nil {
			fmt.Println(w)
			panic(err)
		}
	}

	fmt.Printf("%.2f s elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
