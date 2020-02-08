// 메모 패키지는 동시성에 안전하지 않은 Func 라는 타입의 메모이제이션 함수를 제공한다

package memo

type Memo struct {
	f Func
	cache map[string]result
}

type Func func(key string)(interface{}, error)

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{
		f: f,
		cache: make(map[string]result),
	}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

