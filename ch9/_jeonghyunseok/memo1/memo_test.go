package memo_test

import (
    "testing"
    memo "github.com/gopl-study/gopl.io/ch9/_jeonghyunseok/memo1"
    "github.com/gopl-study/gopl.io/ch9/_jeonghyunseok/memotest"
)

var httpGetBody = memotest.HTTPGetBody 

func TestConcurrency(t *testing.T) {
    m := memo.New(httpGetBody)
    memotest.Concurrent(t, m)
}

func Test(t *testing.T) {
    m := memo.New(httpGetBody)
    memotest.Sequential(t, m)
}



/* 
go test . -v 

하나씩 따라가보자. 

1. httpGetBody 는 결국 memotest 에 있는 함수이다. 
    - url 을 넣으면 바디를 읽어서 결국은 []byte, error 를 리턴해준다.
    - return 값은 interface{} 이므로 .([]byte) 로 타입 어써션 해줘야 할거다
2. memo.New() 를 하면
    1) memo 는 구조체인데 Func 와 cache 라는 map 을 가진다
        - cache 는 string 에 대해 result 라는 구조체를 가진다. 
        - 즉, cache 에 key 값을 가진다면 이미 result 를 가지고 있는 거다. 
            - key 가 없으면 Func 로 가져와서 cache 에 쓸거다
    2) Func 는 string 을 집어넣으면, interface{} 와 error 가 나오는 함수이며
    3) New 는 바로 그 Func 값을 정해서 Memo 를 하나 만들어 내는 것이다
3. 그리고 *Memo 자체적으로도 Get 메쏘드가 정의되어 있다

이제 본격적으로 테스트를 보자
1. 첫번째 Test
    - 함수는 두번째 것도 똑같은 httpGetBody 이다 
    - memotest.Sequential 가 실행된다. 
    - 보면 순서대로 하나씩 m.Get(url) 이 실행되는 것을 알 수 있다.
    - 보면 같은 url 들을 두번 호출하고 있다.

2. 두번째 Test 
    - url 들마다 바로바로 goroutine 으로 처리하고 있다. 
    - 동시성이긴 하다

이제 결과를 보자
1. 두 테스트틑 처음에는 시간이 많이 걸리고, 두번째는 적게 걸린다. 
2. 두번째 테스트가 좀더 빠른거 같긴 하지만 해당 website 자체의 캐시 때문일수도 있다.
    - 이건 반대 순서로 테스트를 해보면 될 것이다.
3. 두 테스트 맨 마지막의 최종 결과를 보면 두번째께 빠르게 나오고 있긴하다. 
    - 암튼 좀 미미한 결과이다.

=== RUN   Test
https://golang.org, 130.6502ms, 11071 bytes
https://godoc.org, 207.4634ms, 7143 bytes
https://play.golang.org, 144.5991ms, 6013 bytes
http://gopl.io, 595.4125ms, 4154 bytes
https://golang.org, 0s, 11071 bytes
https://godoc.org, 0s, 7143 bytes
https://play.golang.org, 0s, 6013 bytes
http://gopl.io, 0s, 4154 bytes
--- PASS: Test (1.08s)
=== RUN   TestConcurrency
https://golang.org, 50.959ms, 11071 bytes
https://godoc.org, 81.8423ms, 7143 bytes
https://play.golang.org, 129.7684ms, 6013 bytes
http://gopl.io, 180.2397ms, 4154 bytes
https://golang.org, 0s, 11071 bytes
https://godoc.org, 0s, 7143 bytes
https://play.golang.org, 0s, 6013 bytes
http://gopl.io, 0s, 4154 bytes
--- PASS: TestConcurrency (0.44s)


순서를 바꾸어서 테스트했더니 TestConcurrency 가 더 오래 걸린다. 
유의미한 테스트를 하려면 url 이 좀 더 많거나 더 오래 걸려야 한다

=== RUN   TestConcurrency
https://golang.org, 123.6688ms, 11071 bytes
https://godoc.org, 120.7147ms, 7143 bytes
https://play.golang.org, 278.2542ms, 6013 bytes
http://gopl.io, 696.1394ms, 4154 bytes
https://golang.org, 0s, 11071 bytes
https://godoc.org, 0s, 7143 bytes
https://play.golang.org, 0s, 6013 bytes
http://gopl.io, 0s, 4154 bytes
--- PASS: TestConcurrency (1.22s)
=== RUN   Test
https://golang.org, 61.8711ms, 11071 bytes
https://godoc.org, 59.8018ms, 7143 bytes
https://play.golang.org, 118.6827ms, 6013 bytes
http://gopl.io, 177.5272ms, 4154 bytes
https://golang.org, 0s, 11071 bytes
https://godoc.org, 0s, 7143 bytes
https://play.golang.org, 0s, 6013 bytes
http://gopl.io, 0s, 4154 bytes
--- PASS: Test (0.42s)

*/