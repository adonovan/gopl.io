package main

import (
	"crypto/sha256"
	"reflect"
	"testing"
)

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func BenchmarkCountSha256Diff(b *testing.B) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	for i := 0; i < b.N; i++ {
		CountSha256Diff(c1, c2)
	}
}

func BenchmarkCountSha256DiffU64(b *testing.B) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	for i := 0; i < b.N; i++ {
		CountSha256DiffU64(c1, c2)
	}
}

func TestCountSha256Diff(t *testing.T) {
	s1 := sha256.Sum256([]byte("x"))
	s2 := sha256.Sum256([]byte("X"))
	s3 := sha256.Sum256([]byte("hello"))

	AssertEqual(t, CountSha256Diff(s1, s1), 0)
	AssertEqual(t, CountSha256Diff(s1, s2), 125)
	AssertEqual(t, CountSha256Diff(s1, s3), 116)
	AssertEqual(t, CountSha256Diff(s2, s3), 133)
	AssertEqual(t, CountSha256Diff(s3, s2), 133)
}

func TestSameCountSha256Diff(t *testing.T) {
	var sample [26][32]byte
	for i, b := range []byte("abcdefghijklmnopqrstuvwxyz") {
		sample[i] = sha256.Sum256([]byte{b})
	}

	for _, a := range sample {
		for _, b := range sample {
			AssertEqual(t, CountSha256Diff(a, b), CountSha256DiffU64(a, b))
		}
	}
}
