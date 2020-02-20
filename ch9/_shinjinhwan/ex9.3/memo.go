package memo

import (
	"fmt"
	"log"
)

// 연습문제 9.3 호출자가 부가적인 done 채널을 통해 작업을 취소(8.9절)할 수 있도록 Func 타입과
// (*Memo).Get 메소드를 확장하라. 취소된 Func 호출의 결과는 캐시하면 안 된다.

// 고루틴과 채널 8장 - 8.9 취소.
// ch8/du4
// var done = make(chan struct{})
//
//func cancelled() bool {
//	select {
//	case <-done:
//		return true
//	default:
//		return false
//	}
//}


type Func func(key string, done <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err error
}

type entry struct {
	res result
	ready chan struct{}
}

type request struct {
	key string
	done <-chan struct{}
	response chan<- result
}

type Memo struct { requests, cancels chan request }

func New(f Func) *Memo {
	// 0. Memo 구조체에 두개의 채널, 요청 & 취소.
	memo := &Memo{ make(chan request), make(chan request)}
	go memo.server(f) // 1.Instance 실행시 실행할 함수 f를 인자로 갖는 Server 고루틴을 생성한다.
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	req := request{key, done,response}
	// send request
	memo.requests <- req
	fmt.Println("get: waiting for response")
	res := <-response
	fmt.Println("get: checking request cancelled")
	// check cancelled request
	select {
		case <-done:
			log.Printf("%s is cancelled", req.key)
			memo.cancels <- req
		default:
	}
	fmt.Println("get: return")
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
Loop:
	for {
	Cancel:
		for {
			select {
			case req := <-memo.cancels:
				fmt.Println("server: deleting cancelled")
				delete(cache, req.key)
			default:
				break Cancel
			}
		}
		select {
		case req := <-memo.cancels:
			fmt.Println("server: deleting cancelled")
			delete(cache, req.key)
			continue Loop
		case req := <-memo.requests:
			fmt.Println("server: request")
			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.done)
			}
			go e.deliver(req.response)
		}
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	e.res.value, e.res.err = f(key, done)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// 4. cache에 키로 등록된 request{} 가 수행되기 까지 대기한다.
	<-e.ready
	response <- e.res
}