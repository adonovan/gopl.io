/* 1.3 잠재적으로 비효율적인 버전과 strings.Join을 사용하는 버전의
  실행 시간 차이를 실험을 통해 측정하라(1.6절에서 time 패키지를 소개하며,
	11.4절은 체계적인 성능 평가를 위한 벤치마크 테스트를 작성하는 방법을 보여준다).
*/
// Use time package this time
// compare echo2, echo3
// reference: https://stackoverflow.com/questions/28755757/time-since-golang-nanosecond-timestamp

package main

import (
	"strings"
	"testing"
)

const sample = "It's too dangerous to go on White Island right now and it's not clear when people will be allowed on Police say it's too dangerous to access White Island right now and according to experts, it's not yet clear when it will be safe.The island remains a no-fly zone, and emergency services are unable to go on the island as there's a risk it could erupt. Dr. Jessica Johnson, a volcanologist at the University of East Anglia in England, said that eruptions could be unpredictable if they involved water.There is a chance of another eruption of similar size ... and potentially bigger, she said. Derek Wyman, a senior lecturer at Sydney University, said the authorities would need to monitor the volcano for around a week to see if the volcanic activity continued, especially as there was no way to predict whether it would erupt again."

var args = strings.Split(sample, " ")

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2(args)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3(args)
	}
}

// go test -bench=.
/*
go test -bench=Echo2
go test -bench=Echo3

go test -bench=. -benchmem  // for cmd
go test -bench . -benchmem  // for PowerShell
*/

// goos: windows
// goarch: amd64
// pkg: gopl.io/presentation/ch1_jeonghyunseok/3-ex1-3
// BenchmarkEcho2-12          86562             13371 ns/op           62824 B/op        144 allocs/op
// BenchmarkEcho3-12         925546              1313 ns/op             896 B/op          1 allocs/op
// PASS
// ok      gopl.io/presentation/ch1_jeonghyunseok/3-ex1-3  2.773s
