package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err) // %v: 원래 형태의 값
			os.Exit(1)                                 // 프로세스가 1의 상태 코드로 종료
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close() // 리소스 노출을 막기 위해
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
