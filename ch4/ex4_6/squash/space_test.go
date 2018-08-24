/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package squash

import (
	"testing"
)

// at a minimum, test for those characters mentioned in unicode.IsSpace:
// '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
func cases() map[int][]string {
	var cases = map[int][]string{}
	cases[0] = []string{
		"a   b     c",
		"a b c",
	}
	cases[1] = []string{
		"a       b",
		"a b",
	}
	cases[2] = []string{
		"a     \n    b",
		"a b",
	}
	cases[3] = []string{
		"a        b       c   ",
		"a b c ",
	}
	cases[4] = []string{
		"a b c\n\td\fe\r\rf",
		"a b c d e f",
	}
	cases[5] = []string{
		"a \u0085 b",
		"a b",
	}
	cases[6] = []string{
		"a \u00a0 b",
		"a b",
	}
	return cases
}

func TestSliceOfBytes_SquashSpaces(t *testing.T) {
	for _, c := range cases() {
		test := c[0]
		want := c[1]
		var sb sliceOfBytes = []byte(test)
		sb.SquashSpaces()
		got := string(sb)
		if got != want {
			t.Errorf("SquashSpaces: got %v, want %v", sb, want)
		}
	}
}
