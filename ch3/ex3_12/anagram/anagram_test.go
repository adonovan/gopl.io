/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */
package anagram

import "testing"

type pair [2]string

var tests = map[pair]bool{
	pair{"a", "b"}:                 false,
	pair{"ab", "ba"}:               true,
	pair{"adsasd1", "12312sdf"}:    false,
	pair{"zookeeper", "oozerpeek"}: true,
	pair{"identity", "identity"}:   false,
}

func TestAnagram(t *testing.T) {
	for test, want := range tests {
		got := Anagram(test[0], test[1])
		if got != want {
			t.Errorf("IsAnagram %s, %s = %v, want %v", test[0], test[1], got, want)
		}
	}
}
