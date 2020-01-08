# 연습문제 4.6

reverse가 UTF-8로 인코딩된 문자열을 나타내는 []byte 슬라이스를 직접 변경해 문자들을 반전시키도록 수정하라.
새 메모리를 할당하지 않고 할 수 있는가?

Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string, in place.
Can you do it without allocating new memory?
<hr>

`rev.go`

```go
func reverse(s []int) { // 파라미터로 slice 를 받는다
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
```