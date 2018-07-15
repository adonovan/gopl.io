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

func TestDedupFast(t *testing.T) {
	for _, c := range cases() {
		in := c[0]
		want := c[1]
		got := DedupFast(in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DedupFast: got %v, want %v", got, want)
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
