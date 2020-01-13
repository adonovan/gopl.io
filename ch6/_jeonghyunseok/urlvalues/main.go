// 이건 뭐하는 놈일까?

package main

import (
	"fmt"
	"net/url"
	"time"
)

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q") == "")
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item") == "")
	time.Sleep(1 * time.Second)
	m.Add("item", "3")

}

// go run main.go

// en
// true
// 1
// [1 2]
// true
// panic: assignment to entry in nil map

// goroutine 1 [running]:
// net/url.Values.Add(...)
//         c:/go/src/net/url/url.go:852
// main.main()
//         d:/golang/src/gopl.io/ch6/_jeonghyunseok/urlvalues/main.go:24 +0x593
// exit status 2
