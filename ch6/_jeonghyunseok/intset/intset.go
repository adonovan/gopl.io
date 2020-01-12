// Package intset 는 비트 벡터에 기반한 정수 set 을 제공한다.package intset
package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

// Add 는 정수를 받아서 64로 나눈 몫과 나머지를 계산한 다음
// 일단 s.words 의 길이보다 크면 s.words 슬라이스에 0 값을 가지는 uint64 를 하나씩 추가한다.
// 즉 슬라이스의 인덱스가 몫 값을 가지는 셈이다.
// 그리고 나서 해당 몫 값을 가지는 index 슬라이스에 해당 나머지 만큼의 비트자리를 1로 만든다
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// 위와 같이 한 상태에서 존재하는지를 확인하려면
// 몫 보다 슬라이스의 길이가 길어야 하고
// 몫 인덱스의 슬라이스 값의 나머지 자리 비트가 1이어야 한다.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// UnionWith 는 두 IntSet 의 합집합(?) 을 만들어준다.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword // s.words 가 감당하는 범위내라면 그냥 | 해준다.
		} else {
			s.words = append(s.words, tword) // 아니면 tword 를 s.words 에 append 해준다
			// 이건 좀 이상하다. 아니구나. t.words 를 순차적으로 증가시키니 문제 없음
		}

	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue // 정수가 없음
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
