// memo2 의 핵심은 Memo 의 cache 에 접근할때 뮤텍스를 사용한다는 것이다.

package memo

import "sync"

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

// 뮤텍스를 이용한 동시성 safe
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}

//!-
