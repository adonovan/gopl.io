/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// TODO: improve test by implementing a "goldenfile"

package wordfreq

import (
	"reflect"
	"strings"
	"testing"
)

type Test struct {
	text string
	want Counts
}

var testTable = []Test{
	{
		"california new\nmexico new hampshire\nnew belgium",
		Counts{
			"california": 1,
			"new":        3,
			"mexico":     1,
			"hampshire":  1,
			"belgium":    1,
		},
	},
}

func TestCount(t *testing.T) {
	for _, test := range testTable {
		t.Run(test.text, func(t *testing.T) {
			in := strings.NewReader(test.text)
			got := Count(in)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
