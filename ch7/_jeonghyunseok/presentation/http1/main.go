// 7.7 The http.Handler Interface 부분 부터이다.

/*
package http
type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}
func ListenAndServe(address string, h Handler) error

1) http.Handler 라는 인터페이스는 ServeHTTP() 라는 메쏘드를 가지고 있으면 되는 거다.
2) 그리고 http.ListenAndServe 라는 함수는 http.Handler 라는 인터페이스 타입을 파라미터로 가진다.
*/

// Http1 is a rudimentary e-commerce server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
이제 코드를 보자.

1) database 라는 map 타입을 선언했는데, 이녀석은 메쏘드로 ServeHTTP() 가 정의되어 있다.
2) 따라서 database 타입은 http.Handler 인터페이스를 만족한다. = http.ListenAndServe() 에 파라미터로 들어갈 수 있다.
3) ServeHTTP() 구현을 보면, w 에 item, price 를 찍어주는게 전부이다.

1) dollars 라는 타입은 String() 이라는 메쏘드를 구현해두었는데, 이게 구현되어 있으면, fmt 에서 %s 타입을 출력할때에 이걸 쓴다.
*/

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

/*
go run main.go
http://localhost:8000

*/
