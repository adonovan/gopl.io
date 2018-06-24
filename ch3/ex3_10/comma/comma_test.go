/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package comma

import (
	"testing"
)

var tests = map[string]string {
	"1" : "1",
	"123" : "123",
	"1024" : "1,024",
	"12345" : "12,345",
	"999888777" : "999,888,777",
	"9129312903812" : "9,129,312,903,812",
}

func TestComma(t *testing.T) {
	for test, want := range tests {
		got := Comma(test)
		if got != want {
			t.Errorf("Comma = %s, want %s", got, want)
		}
	}
}

