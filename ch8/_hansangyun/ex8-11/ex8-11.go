package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/*
8.4.4절의 mirroredQuery 접근 방법에 따라 여러 URL로 동시에 요청을 보내는 fetch의 변형을 구현하라.
첫 번째 응답이 도착하는 대로 다른 요청을 취소하라.
*/

// python -m SimpleHTTPServer 8080
// python3 -m http.server 8080
func main() {
	urls := []string{"https://daum.net", "https://www.naver.com" , "http://localhost:8080"}

	// 캔슬을 위한 채널
	cancel := make(chan struct{})
	responses := make(chan *http.Response)

	wg := &sync.WaitGroup{}
	for _, url := range urls{
		wg.Add(1)

		// URL별로 고루틴을 생성해서
		go func(url string) {
			fmt.Println("====== GOROUTINE ", url," ======")
			defer wg.Done()
			req, err := http.NewRequest("HEAD", url, nil)
			if err != nil {
				log.Printf("[NewRequest] HEAD %s: %s", url, err)
				return
			}
			req.Cancel = cancel
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("[DefaultClient] HEAD %s: %s", url, err)
				return
			}
			responses <- resp
		}(url)
	}

	// 원래는 urls에 요청보낸 갯수만큼 request를 받고 실행해야 하지만
	// 아무거나 하나만 응답
	resp := <-responses
	defer resp.Body.Close()

	// 아직 응답이 오지 않는 요청 취소
	close(cancel)
	fmt.Println(resp.Request.URL)


	wg.Wait()
}