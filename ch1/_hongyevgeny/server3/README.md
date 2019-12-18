## 예제 코드

```
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

	fmt.Fprintf(w, "HOST =  %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
```

## 실행 결과

```
GET / HTTP/1.1
Header["User-Agent"] = ["curl/7.58.0"]
Header["Accept"] = ["*/*"]
HOST =  "localhost:8000"
RemoteAddr = "127.0.0.1:51658"

GET /?help=noThanks HTTP/1.1
Header["User-Agent"] = ["curl/7.58.0"]
Header["Accept"] = ["*/*"]
HOST =  "localhost:8000"
RemoteAddr = "127.0.0.1:51659"
Form["help"] = ["noThanks"]
```
