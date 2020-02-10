package format

import (
	"fmt"
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))
	fmt.Println(Any(d))
	fmt.Println(Any([]int64{x}))
	fmt.Println(Any([]time.Duration{d}))
}

/*
d:\gopath\src\github.com\gopl-study\gopl.io\ch12\_jeonghyunseok\presentation\2.format>go test . -v
=== RUN   TestAny
1
1
[]int64 0xc0000101d0
[]time.Duration 0xc0000101d8
--- PASS: TestAny (0.00s)
PASS */
