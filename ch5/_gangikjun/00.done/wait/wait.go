package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := ""
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}

// WaitForServer 는 url롤 지정된 서버로 접속을 시도한다
// 1분간 지수 단위로 백오프(exponential back-off)를 수행한다
// 모든 시도가 실패하면 오류를 보고한다
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)

		fmt.Printf("%b\n", time.Second)
		fmt.Println(time.Second)

		fmt.Printf("%b\n", time.Second<<uint(tries))
		fmt.Println(time.Second << uint(tries))

		time.Sleep(time.Second << uint(tries)) // 지수 단위 백오프
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
