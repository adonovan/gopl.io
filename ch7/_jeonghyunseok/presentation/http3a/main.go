/*
1) 3에서 본 http.HandlerFunc() 로의 타입 캐스팅은 너무 자주 일어나는 것이니
2) 아예 전담하는 mux.HandleFunc() 라는 넘을 만들어 두었다. a 는 mux 를 하나 만들어서 쓴다
3) mux.HandleFunc() 를 쓰면 좀더 깔끔하게 endpoint 별 handler 를 지정해서 구현할 수 있다.

*/

// Http3a is an e-commerce server that registers the /list and /price
// endpoints by calling (*http.ServeMux).HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	//!+main
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	//!-main
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]int

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: $%d\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "$%d\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
