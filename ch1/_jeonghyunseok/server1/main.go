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

// GET / HTTP/1.1
// Header["Upgrade-Insecure-Requests"] = ["1"]
// Header["Sec-Fetch-Mode"] = ["navigate"]
// Header["Connection"] = ["keep-alive"]
// Header["User-Agent"] = ["Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"]
// Header["Sec-Fetch-User"] = ["?1"]
// Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"]
// Header["Sec-Fetch-Site"] = ["none"]
// Header["Accept-Encoding"] = ["gzip, deflate, br"]
// Header["Accept-Language"] = ["ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7"]
// Host = "localhost:8000"
// RemoteAddr = "127.0.0.1:14964"
