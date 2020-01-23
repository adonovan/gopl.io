// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// go test -bench=. gopl.io/ch8/_jeonghyunseok/cake

package cake_test

import (
	"testing"
	"time"

	"gopl.io/ch8/_jeonghyunseok/cake"
)

/*
Benchmark 는 그냥 디폴트로 돌린다.
BenchmarkBuffers 는 버퍼를 10개씩 둔다.
BenchmarkVariable 는 빵굽는, 식히는, 장식 시간의 표준편차를 1/4 로 줄인다. 작을수록 편차적다
BenchmarkVariableBuffers 는 버퍼 10개 표준편차 1/4 (즉 위 두개 합친거)
BenchmarkSlowIcing 은 원래 0초인데 50 ms 으로 늘린다. 대신 Icer 는 5명
BenchmarkSlowIcingManyIcers 는 위와 같은데 5명이 아니라 50명
*/

func Benchmark(b *testing.B) {
	// Baseline: one baker, one icer, one inscriber.
	// Each step takes exactly 10ms.  No buffers.
	var defaults = cake.Shop{
		Verbose: testing.Verbose(),
		// Verbose:      false,
		Cakes:        20,
		BakeTime:     10 * time.Millisecond,
		NumIcers:     1,
		IceTime:      10 * time.Millisecond,
		InscribeTime: 10 * time.Millisecond,
	}

	cakeshop := defaults
	cakeshop.Work(b.N) // 224 ms
}

func BenchmarkBuffers(b *testing.B) {
	// Adding buffers has no effect.
	var defaults = cake.Shop{
		Verbose: testing.Verbose(),
		// Verbose:      false,
		Cakes:        20,
		BakeTime:     10 * time.Millisecond,
		NumIcers:     1,
		IceTime:      10 * time.Millisecond,
		InscribeTime: 10 * time.Millisecond,
	}

	cakeshop := defaults
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 224 ms
}

func BenchmarkVariable(b *testing.B) {
	// Adding variability to rate of each step
	// increases total time due to channel delays.
	var defaults = cake.Shop{
		Verbose: testing.Verbose(),
		// Verbose:      false,
		Cakes:        20,
		BakeTime:     10 * time.Millisecond,
		NumIcers:     1,
		IceTime:      10 * time.Millisecond,
		InscribeTime: 10 * time.Millisecond,
	}
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.Work(b.N) // 259 ms
}

func BenchmarkVariableBuffers(b *testing.B) {
	// Adding channel buffers reduces
	// delays resulting from variability.
	var defaults = cake.Shop{
		Verbose: testing.Verbose(),
		// Verbose:      false,
		Cakes:        20,
		BakeTime:     10 * time.Millisecond,
		NumIcers:     1,
		IceTime:      10 * time.Millisecond,
		InscribeTime: 10 * time.Millisecond,
	}
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 244 ms
}

func BenchmarkSlowIcing(b *testing.B) {
	// Making the middle stage slower
	// adds directly to the critical path.
	var defaults = cake.Shop{
		Verbose: testing.Verbose(),
		// Verbose:      false,
		Cakes:        20,
		BakeTime:     10 * time.Millisecond,
		NumIcers:     1,
		IceTime:      10 * time.Millisecond,
		InscribeTime: 10 * time.Millisecond,
	}
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.Work(b.N) // 1.032 s
}

func BenchmarkSlowIcingManyIcers(b *testing.B) {
	// Adding more icing cooks reduces the cost of icing
	// to its sequential component, following Amdahl's Law.

	var defaults = cake.Shop{
		Verbose: testing.Verbose(),
		// Verbose:      false,
		Cakes:        20,
		BakeTime:     10 * time.Millisecond,
		NumIcers:     1,
		IceTime:      10 * time.Millisecond,
		InscribeTime: 10 * time.Millisecond,
	}
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.NumIcers = 5
	cakeshop.Work(b.N) // 288ms
}

/*
1. Benchmark 는 그냥 디폴트로 돌린다.
2. BenchmarkBuffers 는 버퍼를 10개씩 둔다.
3. BenchmarkVariable 는 빵굽는, 식히는, 장식 시간의 표준편차를 1/4 로 줄인다. 작을수록 편차적다
4. BenchmarkVariableBuffers 는 버퍼 10개 표준편차 1/4 (즉 위 두개 합친거)
5. BenchmarkSlowIcing 은 원래 0초인데 50 ms 으로 늘린다. 대신 Icer 는 5명
6. BenchmarkSlowIcingManyIcers 는 위와 같은데 5명이 아니라 50명
*/

// d:\golang\src\gopl.io\ch8\_jeonghyunseok\cake>go test -bench=. gopl.io/ch8/_jeonghyunseok/cake
// goos: windows
// goarch: amd64
// pkg: gopl.io/ch8/_jeonghyunseok/cake
// Benchmark-12                                   5         239365680 ns/op
// BenchmarkBuffers-12                            5         240661460 ns/op
// BenchmarkVariable-12                           4         273659525 ns/op
// BenchmarkVariableBuffers-12                    4         251078725 ns/op
// BenchmarkSlowIcing-12                          1        1040251800 ns/op
// BenchmarkSlowIcingManyIcers-12                 4         280256975 ns/op
// PASS
// ok      gopl.io/ch8/_jeonghyunseok/cake 12.121s

/*
1. 디폴트로 돌리면 239365680 ns 마다 한번 돌린다
2. 버퍼를 10개씩 둬도 큰 차이가 없다.
3. 표준편차를 줄였는데 오히려 시간이 늘었다. 이건 그때그때 다를듯하다.
4. 2, 3 합친건데 희한하게 이건 더 안좋다.
5. Icer 를 늘리긴 했는데 icing 시간이 늘었다. 최악이다.
6. 그래도 Icer 를 왕창 늘리니깐 시간이 줄어들긴 한다.

*/
