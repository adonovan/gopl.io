// 4는 뭐가 다를까? 
package memo

import "sync"

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// 여기도 달라졌구나. 
// resource 준비여부를 챙기는 부분이다. 
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// 여기서부터 달라짐. 메모 struct 내부에 뮤텍스가 들어감
type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}


// 여기도 많이 달라졌다.
// 1. 일단 락을 걸고
// 2. 값이 없을때만 == function 을 실행해서 값을 가져오고, 채워넣어야 할때만 여기를 돈다.
	// - 여기서 여러 놈들이 동시게 function을 돌리려 한다 생각해보자
	// - 맨 첫번째 놈이 e 에 값을 채워넣는다. 이제 nil 아님. 흠 그런데 여기까지 동시에 발생?
	// 	- 동시에 진입할 가능성이 없진 않겠다. 
	// - 그러면 나머지 놈들은 else 로 가고 <-e.ready 가 close() 되기를 기다린다. 
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}
/* 

go test . -v
=== RUN   Test
https://golang.org, 309.1791ms, 11071 bytes
https://godoc.org, 151.5956ms, 7143 bytes
https://play.golang.org, 162.558ms, 6013 bytes
http://gopl.io, 631.3149ms, 4154 bytes
https://golang.org, 0s, 11071 bytes
https://godoc.org, 0s, 7143 bytes
https://play.golang.org, 0s, 6013 bytes
http://gopl.io, 0s, 4154 bytes
--- PASS: Test (1.25s)
=== RUN   TestConcurrent
https://golang.org, 42.9262ms, 11071 bytes
https://godoc.org, 58.112ms, 7143 bytes
https://play.golang.org, 119.6849ms, 6013 bytes
http://gopl.io, 198.5619ms, 4154 bytes
https://golang.org, 0s, 11071 bytes
https://godoc.org, 0s, 7143 bytes
https://play.golang.org, 0s, 6013 bytes
http://gopl.io, 0s, 4154 bytes
--- PASS: TestConcurrent (0.42s)
PASS */