// Wait 은 HTTP server 가 resp 를 시작하는 걸 기다린다.package wait
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// WaitForServer 는 30 초동안 계속 http.Head() 를 시도한다.
// 시도 간격은 1, 2, 4, 8, ... 초
// resp 를 받으면 성공. 끝내 못받으면 실패 리턴

// HTTP HEAD 메소드는 특정 리소스를 HTTP GET 메소드로 요청하는 경우에
// 어떤 헤더들이 반환되는지를 요청합니다. 예를 들어, 큰 용량의 리소스를 다운로드 받을지
// 말지 결정하기 위해서 사전 요청하는 용도로 사용할 수 있습니다.

func WaitForServer(url string) error {
	const timeout = 30 * time.Second
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url) // Head 는 뭐지?
		if err == nil {
			log.Println("get response: url", url)
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: wait url...\n")
	}

	for _, url := range os.Args[1:] {
		if err := WaitForServer(url); err != nil {
			log.Fatalf("Site URL is wrong, or Site is down: %v", err)
		}
	}
}

/*
go build -o wait.exe
wait.exe http://www.google.com http://www.dfjliwjonv.co.kr

*/

// d:\golang\src\gopl.io\ch5\_jeonghyunseok\wait>wait.exe http://www.google.com http://www.dfjliwjonv.co.kr
// 2020/01/08 18:41:24 get response: url http://www.google.com
// 2020/01/08 18:41:24 server not responding (Head http://www.dfjliwjonv.co.kr: dial tcp: lookup www.dfjliwjonv.co.kr: no such host); retrying...
// 2020/01/08 18:41:25 server not responding (Head http://www.dfjliwjonv.co.kr: dial tcp: lookup www.dfjliwjonv.co.kr: no such host); retrying...
// 2020/01/08 18:41:27 server not responding (Head http://www.dfjliwjonv.co.kr: dial tcp: lookup www.dfjliwjonv.co.kr: no such host); retrying...
// 2020/01/08 18:41:31 server not responding (Head http://www.dfjliwjonv.co.kr: dial tcp: lookup www.dfjliwjonv.co.kr: no such host); retrying...
// 2020/01/08 18:41:39 server not responding (Head http://www.dfjliwjonv.co.kr: dial tcp: lookup www.dfjliwjonv.co.kr: no such host); retrying...
// 2020/01/08 18:41:55 Site URL is wrong, or Site is down: server http://www.dfjliwjonv.co.kr failed to respond after 30s
