// 이건 앞에꺼와 똑같다.
package memo_test

import (
	"testing"

	memo "github.com/gopl-study/gopl.io/ch9/_jeonghyunseok/memo2"
	"github.com/gopl-study/gopl.io/ch9/_jeonghyunseok/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}

/*
go test . -v

다좋은데 시퀸셜과 동시성의 차이점이 없어져버렸다.
결국 뮤텍스 앞에서 줄서서 기다려야 하니깐
=== RUN   Test
https://golang.org, 203.4555ms, 11071 bytes
https://godoc.org, 131.6478ms, 7143 bytes
https://play.golang.org, 154.5866ms, 6013 bytes
http://gopl.io, 562.5351ms, 4154 bytes
https://golang.org, 0s, 11071 bytes
https://godoc.org, 0s, 7143 bytes
https://play.golang.org, 0s, 6013 bytes
http://gopl.io, 0s, 4154 bytes
--- PASS: Test (1.05s)
=== RUN   TestConcurrent
https://golang.org, 40.8498ms, 11071 bytes
https://godoc.org, 138.6293ms, 7143 bytes
https://play.golang.org, 119.6797ms, 6013 bytes
http://gopl.io, 175.5306ms, 4154 bytes
https://golang.org, 0s, 11071 bytes
https://godoc.org, 0s, 7143 bytes
https://play.golang.org, 0s, 6013 bytes
http://gopl.io, 0s, 4154 bytes
--- PASS: TestConcurrent (0.47s)
PASS
*/
