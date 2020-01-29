/*
인상적인 부분 하나 보고 가자.

Go doesn’t have a canonical web framework analogous to
Ruby’s Rails or Python’s Django. This is not to say that such frameworks don’t exist, but the
building blocks in Go’s standard library are flexible enough that frameworks are often
unnecessary. Furthermore, although frameworks are convenient in the early phases of a
project, their additional complexity can make longer-term maintenance harder.
*/

/*
1) 3은 mux 를 하나 만들어서 쓰는 거다.

2)http1 은 http.Handler 인터페이스를 만족하는 타입을 하나 만들어서
	http.ListenAndServe() 의 파라미터로 넣었었다.
	여기서는 *http.ServeMux 타입을 넣어준다.

*http.ServeMux 타입을 찾아 들어가보면 ServeHTTP() 가 구현되어 있다.
따라서 http.ListenAndServe() 의 파라미터로 넣을 수 있는 거다.

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}

3) mux.Handle() 에서 endpoint 와 실제 핸들러를 연결시켜주는 게 보인다.
4) http.HandlerFunc() 는 함수가 아니다. int(1.0) 처럼 타입 캐스팅이다.
	- 일반적인 함수를 HTTP handler 함수로 쓸 수 있게 해주는 거다.
	- 1) 일반 함수를 HandlerFunc 함수로 타입을 바꿔주면
	- 2) HandlerFunc 함수는 ServeHTTP() 메쏘드를 가지고 있으니 ListenAndServe() 에 쓸 수 있게 되는 것임

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

*/

// Http3 is an e-commerce server that registers the /list and /price
// endpoints by calling (*http.ServeMux).Handle.
package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

//!-main

/*
//!+handlerfunc
package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
//!-handlerfunc
*/
