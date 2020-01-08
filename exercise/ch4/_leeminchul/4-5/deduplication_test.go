package main

import (
	"fmt"
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

func TestCorrectResult(t *testing.T) {
	AssertEqual(
		t,
		deduplicate([]string{"a", "a", "b", "c", "c", "d", "e", "f", "g", "g"}),
		[]string{"a", "b", "c", "d", "e", "f", "g"},
	)

	AssertEqual(
		t,
		deduplicate([]string{"a", "b", "c", "d", "e"}),
		[]string{"a", "b", "c", "d", "e"},
	)

	AssertEqual(
		t,
		deduplicate([]string{"a", "b", "c", "c", "d", "d", "d", "e"}),
		[]string{"a", "b", "c", "d", "e"},
	)
}

func TestSamePointer(t *testing.T) {
	before := []string{"a", "a", "b", "c", "c", "d", "e", "f", "g", "g"}
	after := deduplicate(before)

	AssertEqual(
		t,
		fmt.Sprintf("%p", before),
		fmt.Sprintf("%p", after),
	)
}

func TestNilSlice(t *testing.T) {
	AssertEqual(
		t,
		deduplicate(make([]string, 0, 0)),
		make([]string, 0, 0),
	)
}

func TestNil(t *testing.T) {
	AssertEqual(
		t,
		deduplicate(nil),
		[]string(nil),
	)
}
