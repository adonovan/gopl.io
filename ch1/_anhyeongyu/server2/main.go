// Server2는 최소한 "echo" 및 카운터 서버다.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

// 핸들러는 요청된 URL의 Path 구성 요소를 반환한다.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // 동시 접근을 차단
	count++
	mu.Unlock() // 동시 접근을 차단
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter는 지금까지 요청된 수를 반환한다.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
