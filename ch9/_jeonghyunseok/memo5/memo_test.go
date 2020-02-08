// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"testing"

	memo "github.com/gopl-study/gopl.io/ch9/_jeonghyunseok/memo5"
	"github.com/gopl-study/gopl.io/ch9/jeonghyunseok/memotest"
)

var httpGetBody = memotest.HTTPGetBody

// 여기서는 왜인지는 몰라도 m.Close() 를 시켜주네. 그럼 메모속의 request 를 close 해준다는 거구나

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}
