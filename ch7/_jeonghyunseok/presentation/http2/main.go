/*
이제 endpoint 를 달리 하면서 거기에 맞게 동작하게 해보자.
req.URL.Path 에 endpoint 정보가 담겨있다.

go run main.go

1) 이걸 출력하게 ServeHTTP() 를 약간 수정해본다. fmt.Println("endpoint:", req.URL.Path)
2) /list 는 http1 에서 구현했던 내용이다.
	http://localhost:8000/list
3) /price 는 HTTP GET 에서 query 를 이용하여 파라미터를 전달해주면 그 값을
	http://localhost:8000/price?item=shoes
	http://localhost:8000/price?item=socks
	http://localhost:8000/price?item=pants
	http://localhost:8000/price?item=shorts
4) 잘못된 endpoint 는 오류 전달
	http://localhost:8000/admin


*/

// Http2 is an e-commerce server with /list and /price endpoints.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5, "pants": 100}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

//!+handler
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("endpoint:", req.URL.Path)
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "price of %s is %s\n", item, price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}
