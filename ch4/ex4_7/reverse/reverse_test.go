/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package reverse

import (
	"testing"
)

func cases() map[int][]string {
	var cases = map[int][]string{}
	cases[0] = []string{
		"tap",
		"pat",
	}
	cases[1] = []string{
		"hero",
		"oreh",
	}
	cases[2] = []string{
		"tacocat",
		"tacocat",
	}
	return cases
}

func TestReverse(t *testing.T) {
	for _, c := range cases() {
		test := []byte(c[0])
		want := []byte(c[1])
		Reverse(test)
		if string(test) != string(want) {
			t.Errorf("Reverse: got %v, want %v", string(test), string(want))
		}
	}
}
