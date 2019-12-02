/* 1.3 잠재적으로 비효율적인 버전과 strings.Join을 사용하는 버전의
  실행 시간 차이를 실험을 통해 측정하라(1.6절에서 time 패키지를 소개하며,
	11.4절은 체계적인 성능 평가를 위한 벤치마크 테스트를 작성하는 방법을 보여준다).
*/
// Use time package this time
// compare echo2, echo3
// reference: https://stackoverflow.com/questions/28755757/time-since-golang-nanosecond-timestamp

package main

import (
	"testing"
)

var args = []string{"Who", "let", "the", "dogs", "out!"}

func BenchmarkEcho2(b *testing.B) {
	echo2(args)
}

func BenchmarkEcho3(b *testing.B) {
	echo3(args)
}

// go test -bench=.
