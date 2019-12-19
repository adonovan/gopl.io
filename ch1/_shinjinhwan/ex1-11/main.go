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

	nbytes, err := io.CopyBuffer(ioutil.Discard, resp.Body, [104]byte)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

/*
./ex1-11 https://Google.com https://Naver.com https://Youtube.com https://Daum.net https://Tmall.com https://Tistory.com https://Google.co.kr https://Facebook.com https://Sohu.com https://Qq.com https://Login.tmall.com https://Taobao.com https://Wikipedia.org https://Namu.wiki https://Amazon.com https://Jd.com https://Kakao.com https://360.cn https://Dcinside.com https://Baidu.com https://Netflix.com https://Coupang.com https://Pages.tmall.com https://Blog.me https://Weibo.com https://Sina.com.cn https://11st.co.kr https://Gmarket.co.kr https://Torrentwal.com https://Apple.com https://Donga.com https://Microsoft.com https://Stackoverflow.com https://Yahoo.com https://Twitch.tv https://Ebay.com https://Inven.co.kr https://Adobe.com https://Auction.co.kr https://Office.com https://Clien.net https://Nexon.com https://Bing.com https://Nate.com https://Amazon.co.jp https://Dropbox.com https://Tumblr.com https://Instagram.com https://Interpark.com https://Amazon.co.uk

Write will be in 'resp-result.txt'
 */
