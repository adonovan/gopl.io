package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}} // 직접 생성
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // en
	fmt.Println(m.Get("q"))    // empty string
	fmt.Println(m.Get("item")) // 1
	fmt.Println(m["item"])     // [1 2]

	m = nil
	fmt.Println(m.Get("item"))
	// m.Add("item", "3") // panic
}
