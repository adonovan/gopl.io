/*
앞의 예제들에서 우리는 mux 를 따로 생성해서 사용했는데
서버가 여러개라면 몰라도 하나뿐이라면 - 사실 하나면 왠만한 환경에서는 충분하지? - mux 하나면 충분하다.
따라서 Go 에서는 기본으로 전역 instance 를 제공한다.

DefaultServeMux 이다. 여기에 endpoint 와 그 Handler를 등록하려면 http.HandleFunc() 를 사용해주면 된다.

	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
		DefaultServeMux.HandleFunc(pattern, handler)
	}

*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

/*
7.11 클라이언트가 데이터베이스 항목을 생성, 조회, 갱신, 삭제할 수 있도록 부가적인 핸들러를 추가하라.
예를 들어 /update?item=socks&price=6 형태의 요청은 재고 목록에서 항목의 값을 갱신하고,
항목이 없거나 가격이 잘못된 경우 오류를 보고한다(경고：이 변경은 변수가 동시에 갱신되게 한다).
-> a) endpoint 만 추가해주면 되겠지?
-> b) Request 마다 goroutine 으로 핸들러가 동작하는데 서버내의 같은 변수들을 참조, 갱신한다면?
-> c) 동시성에 대한 고민이 필요한거다

7.12 /list 핸들러를 변경해 텍스트가 아닌 HTML 테이블로 출력하라. html/template 패키지(4.6절)가 유용할 것이다.
-> a) 이런 query 를 날려주는 버튼 하나 만들고
-> b) 결과를 HTML 로 출력하게 해주면 웹서버 인거다.
*/
