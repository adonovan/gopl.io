// Server2 is
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
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

/*
go build -o server2.exe
server2.exe

on browser
http://localhost:8000/
http://localhost:8000/abcd
http://localhost:8000/count
http://localhost:8000/count
http://localhost:8000/count
*/

// URL.Path = "/"
// URL.Paht = "/abcd"
// Count 2
// Count 2
// Count 2
