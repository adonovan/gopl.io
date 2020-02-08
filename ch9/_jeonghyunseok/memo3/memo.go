// 메모3은 뭐가 다를까?
// 일단 나머지는 같고, 메쏘드 Get 만 다르다
package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}


// 이번에는 memo.cache 에 접근할때만 락을 걸게 해두었다. 
// 하지만 완벽해보이지는 않는다. 계산하는 memo.f(key) 가 동시에 여러개 돌아갈 가능성이 있다.
// 아무튼 지저분하다. 여기저기 뮤텍스
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// Between the two critical sections, several goroutines
		// may race to compute f(key) and update the map.
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}


