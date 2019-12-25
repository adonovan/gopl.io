// Server1은 최소한의 "echo" 서버다.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)                         // 각 요청은 핸들러를 호출한다.
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // log.Fatal(): 메시지를 출력하고 os.Exit(1)을 호출하여 프로그램을 종료
}

// handler는 요청된 URL r의 Path 구성 요소를 반환한다.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path) // %q: 따옴표로 묶인 문자열 "abc" 또는 룬 'c'
}
