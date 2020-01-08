package main

import (
	"fmt"
	"reflect"
	"testing"
)

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, a []byte, b []byte) {
	if reflect.DeepEqual(a, b) {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", string(a), reflect.TypeOf(a), string(b), reflect.TypeOf(b))
}

func TestCorrectResult(t *testing.T) {
	AssertEqual(
		t,
		mergeSpace([]byte("독도는우리땅")),
		[]byte("독도는우리땅"),
	)

	AssertEqual(
		t,
		mergeSpace([]byte(" 독도는 우리 땅")),
		[]byte(" 독도는 우리 땅"),
	)

	AssertEqual(
		t,
		mergeSpace([]byte("  독\xc2\x85도 \xc2\x85는\xc2\x85 우리\xc2\xa0\xc2\x85땅   ")),
		[]byte(" 독 도 는 우리 땅 "),
	)
}

func TestSamePointer(t *testing.T) {
	before := []byte("  독\xc2\x85도 \xc2\x85는\xc2\x85 우리\xc2\xa0\xc2\x85땅   ")
	after := mergeSpace(before)

	if fmt.Sprintf("%p", before) != fmt.Sprintf("%p", after) {
		t.Errorf("%p not equals %p", before, after)
	}
}

func TestNilSlice(t *testing.T) {
	AssertEqual(
		t,
		mergeSpace(make([]byte, 0, 0)),
		make([]byte, 0, 0),
	)
}

func TestNil(t *testing.T) {
	AssertEqual(
		t,
		mergeSpace(nil),
		[]byte(nil),
	)
}
