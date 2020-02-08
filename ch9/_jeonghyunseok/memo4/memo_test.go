// 이건 바뀐게 없다.
package memo_test

import (
	"testing"

	memo "github.com/gopl-study/gopl.io/ch9/_jeonghyunseok/memo4"
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
