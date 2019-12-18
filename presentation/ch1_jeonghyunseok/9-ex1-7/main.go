// 1.7 io.Copy(dst, src) 함수 호출은 src에서 읽어 dst에 기록한다.
// ioutil.ReadAll 대신 이 함수를 사용해 결과 본문을 전체 스트림을 저장하는 큰 버퍼 없이
// os.Stdout으로 복사하라.
// io.Copy의 오류 결과를 확인해야 한다.

// start with fetch.go
// ref link: https://haisum.github.io/2017/09/11/golang-ioutil-readall/

// fetch is
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("ex1-7.exe")
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)  // using whole memory for byte slice
		io.Copy(os.Stdout, resp.Body) //  using 32KB buffer
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// fmt.Printf("%s", b)
	}
}

/*
go build -o ex1-7.exe
ex1-7.exe http://www.president.go.kr/
ex1-7.exe http://www.bad.gopl.io

*/
