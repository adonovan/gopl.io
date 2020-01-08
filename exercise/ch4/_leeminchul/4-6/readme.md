# 연습문제 4.6

UTF-8로 인코딩된 []byte 슬라이스 안의 인접한 유니코드 공백(`unicode.IsSpace`를 보라)을 단일 아스키 공백으로 합치는 직접 변경 함수를 작성하라.

Write an in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.

<hr>

*unicode.IsSpace*
```go
func IsSpace(r rune) bool {
	// This property isn't the same as Z; special-case it.
	if uint32(r) <= MaxLatin1 {
		switch r {
		case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
			return true
		}
		return false
	}
	return isExcludingLatin(White_Space, r)
}
```

-> 인접한 (`'\t'`, `'\n'`, `'\v'`, `'\f'`, `'\r'`, `' '`, `0x85`, `0xA0`) 를 `' '` 로 합쳐라.  
