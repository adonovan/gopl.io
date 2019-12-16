// Server 1
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

/*
go build -o server1.exe
server1.exe

on browser
http://localhost:8000/
http://localhost:8000/abcd

*/

// URL.Path = "/"
// URL.Path = "/abcd"
