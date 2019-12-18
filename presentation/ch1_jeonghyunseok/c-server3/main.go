// Server 3 is

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
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q]\n", k, v)
	}
}

/*
go build -o server3.exe
server3.exe

on web
http://localhost:8000
*/

// GET / HTTP/1.1
// Header["Accept-Language"] = ["ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7"]
// Header["User-Agent"] = ["Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"]
// Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"]
// Header["Sec-Fetch-Mode"] = ["navigate"]
// Header["Accept-Encoding"] = ["gzip, deflate, br"]
// Header["Cookie"] = ["session_token=797314ff-79ea-4463-96ba-4a0292fa314c"]
// Header["Connection"] = ["keep-alive"]
// Header["Upgrade-Insecure-Requests"] = ["1"]
// Header["Sec-Fetch-User"] = ["?1"]
// Header["Sec-Fetch-Site"] = ["none"]
// Host = "localhost:8000"
// RemoteAddr = "127.0.0.1:4693"
