// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount

import (
	"log"
	"testing"
)

func TestPopCount(t *testing.T) {
	// this is the gopl book function on p.45
	want := PopCount(0x1234567890ABCDEF)
	// this is my looped implementation
	got := PopCountWithLoop(0x1234567890ABCDEF)
	if want != got {
		log.Fatalf("PopCountWithLoop: got %v, want %v", got, want)
	}
}

func TestPopCountDeadbeefs(t *testing.T) {
	// this is the gopl book function on p.45
	want := PopCount(0xdeadbeef)
	// this is my looped implementation
	got := PopCountWithLoop(0xdeadbeef)
	if want != got {
		log.Fatalf("PopCountWithLoop: got %v, want %v", got, want)
	}
}

func TestPopCountCoffeeDeadbeef(t *testing.T) {
	want := PopCount(0xc0ffee)
	got := PopCountWithLoop(0xdeadbeef)
	if want == got {
		log.Fatalf("PopCount: unexpected result")
	}
}

func TestPopCountWithLoop(t *testing.T) {
	PopCountWithLoop(0x1234567890ABCDEF)
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithLoop(0x1234567890ABCDEF)
	}
}
