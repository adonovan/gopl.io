// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"testing"

	memo "github.com/gopl-study/gopl.io/ch9/_jeonghyunseok/memo3"
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
테스트시간이 엄청 늘어났다.

=== RUN   Test
https://golang.org, 485.7137ms, 11071 bytes
https://godoc.org, 592.4008ms, 7143 bytes
https://play.golang.org, 189.5169ms, 6013 bytes
http://gopl.io, 2.6449097s, 4154 bytes
https://golang.org, 0s, 11071 bytes
https://godoc.org, 0s, 7143 bytes
https://play.golang.org, 0s, 6013 bytes
http://gopl.io, 0s, 4154 bytes
--- PASS: Test (3.91s)
=== RUN   TestConcurrent
https://golang.org, 800.8657ms, 11071 bytes
https://godoc.org, 571.4378ms, 7143 bytes
https://play.golang.org, 366.0425ms, 6013 bytes
http://gopl.io, 2.0604754s, 4154 bytes
https://golang.org, 0s, 11071 bytes
https://godoc.org, 0s, 7143 bytes
https://play.golang.org, 0s, 6013 bytes
http://gopl.io, 0s, 4154 bytes
--- PASS: TestConcurrent (3.80s)
PASS

*/
