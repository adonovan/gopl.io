/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package charcount

import (
	"testing"
)

type CharcountTest struct {
	Input string
	Want  Counts
}

var tests = []CharcountTest{
	{
		"hi", Counts{
			characters: 2,
			graphic:    2,
			letter:     2,
			lowerCase:  2,
		},
	}, {
		"hi5", Counts{
			characters: 3,
			digit:      1,
			number:     1,
			graphic:    3,
			letter:     2,
			lowerCase:  2,
		},
	}, {
		"Fnord!", Counts{
			characters:  6,
			graphic:     6,
			letter:      5,
			lowerCase:   4,
			upperCase:   1,
			punctuation: 1,
		},
	},
	{
		"\b", Counts{
			characters: 1,
			control:    1,
		},
	}, {
		"℃", Counts{
			characters: 1,
			graphic:    1,
			symbol:     1,
		},
	},
}

func TestCharcount(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.Input, func(t *testing.T) {
			got := Charcount(tc.Input)
			if got != tc.Want {
				t.Errorf("Charcount: for %s\ngot:\n%+v\nwant:\n%+v",
					tc.Input,
					got,
					tc.Want,
				)
			}
		})
	}
}
