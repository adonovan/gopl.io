/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package charcount

import (
	"testing"
)

var cases = map[string]Counts{}

func init() {
	cases["hi"] = Counts{
		characters: 2,
		graphic:    2,
		letter:     2,
		lowerCase:  2,
	}
	cases["hi5"] = Counts{
		characters: 3,
		digit:      1,
		number:     1,
		graphic:    3,
		letter:     2,
		lowerCase:  2,
	}
	cases["Fnord!"] = Counts{
		characters:  6,
		graphic:     6,
		letter:      5,
		lowerCase:   4,
		upperCase:   1,
		punctuation: 1,
	}
	cases["\b"] = Counts{
		characters: 1,
		control:    1,
	}
	cases["℃"] = Counts{
		characters: 1,
		graphic:    1,
		symbol:     1,
	}
}

func TestCharcount(t *testing.T) {
	for test, want := range cases {
		got := Charcount(test)
		if got != want {
			t.Errorf("Charcount: for %s\ngot:\n%+v\nwant:\n%+v", test, got, want)
		}
	}
}
