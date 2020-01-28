package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}} //직접 생성
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1" (첫 번째 값)
	fmt.Println(m["item"])     // "[1 2]" (직접 맵 접근)

	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // 패닉: nil 맵의 원소로 할당
}
