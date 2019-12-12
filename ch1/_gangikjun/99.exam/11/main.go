// Fetchall은 url을 병렬로 반입하고 시간과 크기를 보고한다
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// https://Google.com https://Youtube.com https://Tmall.com https://Facebook.com https://Baidu.com https://Qq.com https://Sohu.com https://Taobao.com https://Login.tmall.com https://Wikipedia.org https://Yahoo.com https://Amazon.com https://Jd.com https://360.cn https://Weibo.com https://Sina.com.cn https://Pages.tmall.com https://Reddit.com https://Live.com https://Vk.com https://Okezone.com https://Netflix.com https://Blogspot.com https://Alipay.com https://Csdn.net https://Aliexpress.com https://Instagram.com https://Office.com https://Yahoo.co.jp https://Xinhuanet.com https://Bing.com https://Microsoft.com https://Stackoverflow.com https://Twitter.com https://Google.com.hk https://Livejasmin.com https://Ebay.com https://Naver.com https://Babytree.com https://Amazon.co.jp https://Twitch.tv https://Tribunnews.com https://Google.co.in https://Apple.com https://Pornhub.com https://Tianya.cn https://Microsoftonline.com https://Msn.com https://Yandex.ru https://Wordpress.com

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
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
