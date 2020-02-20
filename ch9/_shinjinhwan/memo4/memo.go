package memo

import "sync"

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// 중복 억제를 위한 구조체.
type entry struct {
	res   result
	ready chan struct{}
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// cache는 각 주소에 대한 entry 구조체의 포인터.
		// 추가로 `ready` 라는 채널이 포함됨.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready)
	} else {
		memo.mu.Unlock()

		<-e.ready
	}
	return e.res.value, e.res.err
}
