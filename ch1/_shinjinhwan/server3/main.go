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
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w,"Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Header {
		fmt.Fprintf(w,"Header[%q] = %q\n", k, v)
	}
}


/*
$ curl localhost:8000

GET / HTTP/1.1
Header["User-Agent"] = ["curl/7.54.0"]
Header["Accept"] = ["*/*"]
Host = "localhost:8000"
RemoteAddr = "127.0.0.1:65357"
Header["User-Agent"] = ["curl/7.54.0"]
Header["Accept"] = ["*/*"]

 */