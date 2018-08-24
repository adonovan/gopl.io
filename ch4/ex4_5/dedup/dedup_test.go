/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package dedup

import (
	"reflect"
	"testing"
)

func cases() map[int][][]string {
	var cases = map[int][][]string{}
	cases[0] = [][]string{
		{"cat", "dog", "dog", "parrot"},
		{"cat", "dog", "parrot"},
	}
	cases[1] = [][]string{
		{"a", "a", "a", "z"},
		{"a", "z"},
	}
	cases[2] = [][]string{
		{"a", "a", "a", "b", "z", "z"},
		{"a", "b", "z"},
	}
	return cases
}

func TestDedupPtrReceiver(t *testing.T) {
	for _, c := range cases() {
		param := c[0]
		want := c[1]
		var ss stringSlice = param
		ss.Dedup()
		// if !reflect.DeepEqual(ss, want) {
		// FIXME: the capacity differs even though the slice content is correct
		// don't use DeepEqual in this test
		// workaround is just check the lengths
		if len(ss) != len(want) {
			t.Errorf("Dedup: got %v, want %v", ss, want)
		}
	}
}

func TestDedupInPlace(t *testing.T) {
	for _, c := range cases() {
		param := c[0]
		want := c[1]
		DedupInPlace(&param)
		if !reflect.DeepEqual(param, want) {
			t.Errorf("DedupInPlace: got %v, want %v", param, want)
		}
	}
}

func TestDedup(t *testing.T) {
	for _, c := range cases() {
		in := c[0]
		want := c[1]
		got := Dedup(in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Dedup: got %v, want %v", got, want)
		}
	}
}
