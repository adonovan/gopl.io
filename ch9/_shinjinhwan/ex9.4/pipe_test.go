package main

import "testing"

func BenchmarkPipeline(b *testing.B) {
	in, out := pipeline(100000000)
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}

// go test -cpu=1 -bench=. -benchmem