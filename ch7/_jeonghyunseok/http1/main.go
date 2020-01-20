// 초보적인 전자상거래 서버이다. rudimentary

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

/*
go build -o http1.exe
http1.exe
http://localhost:8000

*/

// 아래는 인터페이스 참고

// package http

// type Handler interface {
// 	ServeHTTP(w ResponseWriter, r *Request)
// }

// func ListenAndServe(address string, h Handler) error