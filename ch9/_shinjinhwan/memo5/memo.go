package memo

type Func func(key string) (interface{}, error)

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
	response chan<- result
}

type Memo struct{ requests chan request }

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f) // 1.Instance 실행시 실행할 함수 f를 인자로 갖는 Server 고루틴을 생성한다.
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	// 2. 함수 f 의 입력값인 `key`와 `response` 채널을 갖는 `request` 구조체를 requests 채널로 보낸다.
	memo.requests <- request{key, response}
	// 5. memo.server 고루틴에서 해당 request에 대한 응답을 받는다.
	// 여기서 요청한 응답이 아닌 이전에 request가 올 수 있다.
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)

	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// 3-1. key 에 해당하는 함수 f를 수행하고 있는 고루틴이 없는 경우
			// 응답을 대기하는 채널을 가지는 entry 구조체를 맵 구조체에 등록(?) 한다.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			// 해당 함수(httpGetBody)와 url을 호출하는 `call` 메소드 고루틴 생성한다.
			go e.call(f, req.key)
		}
		// 3-2. cache 맵에 key 에 해당하는 함수 f 를 동작하는 고루틴이 있는경우,
		// `deliver` 메소드를 통해 대기하는 고루틴을 생성한다.
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
}

func (e *entry) deliver(response chan<- result) {
	// 4. cache에 키로 등록된 request{} 가 수행되기 까지 대기한다.
	<-e.ready
	response <- e.res
}