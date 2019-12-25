# net/http

## HandleFunc

핸들러 로직을 만들 때 사용하는 함수

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

만약 "/api" url에 대한 핸들러 함수를 등록하고 싶다면,

```go
http.HandleFunc("/api", func(w http.ResponseWriter, r * http.Request) {
    // 핸들러 로직 작성...
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
})

http.ListenAndServe("localhost:8000", nil)
```

## ListenAndServe

요청 대기 상태인 서버를 구동하는 함수

```go
func ListenAndServe(addr string, handler Handler) error
```
