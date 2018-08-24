package ex1_3

import (
	"testing"
)

func TestEcho1(t *testing.T) {
	input := []string{"cmd", "a", "b", "c"}
	want := "a b c"
	got := echo1(input)
	if got != want {
		t.Errorf("echo1(%v) = %v, want %v", input, got, want)
	}
}

func BenchmarkEcho1(b *testing.B) {
	input := []string{"cmd", "a", "b", "c"}
	for i := 0; i < b.N; i++ {
		echo1(input)
	}
}

func TestEcho3(t *testing.T) {
	input := []string{"cmd", "a", "b", "c"}
	want := "a b c"
	got := echo3(input)
	if got != want {
		t.Errorf("echo1(%v) = %v, want %v", input, got, want)
	}
}

func BenchmarkEcho3(b *testing.B) {
	input := []string{"cmd", "a", "b", "c"}
	for i := 0; i < b.N; i++ {
		echo3(input)
	}
}
