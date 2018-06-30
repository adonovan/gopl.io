/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package comma

import (
	"bytes"
	"regexp"
)

var signRegexp = regexp.MustCompile(`\+|\-`)
var decimalRegexp = regexp.MustCompile(`\.\d+`)

// Comma inserts commas at each 1000's place in a numeric string. It allows
// a leading + or - sign, and trailing decimal point digits.
func Comma(s string) string {
	buf := bytes.Buffer{}
	// default to formatting entire string
	startIndex := 0
	endIndex := len(s)
	// check for sign
	signIndex := signRegexp.FindStringIndex(s)
	// check for decimal point
	decimalIndex := decimalRegexp.FindStringIndex(s)
	if signIndex != nil {
		startIndex = signIndex[0] + signIndex[1]
	}
	if decimalIndex != nil {
		endIndex = decimalIndex[0]
	}
	substring := s[startIndex:endIndex]
	formatted := formatWithCommas(substring)
	if signIndex != nil {
		buf.WriteString(s[signIndex[0]:signIndex[1]])
	}
	buf.WriteString(formatted)
	if decimalIndex != nil {
		buf.WriteString(s[decimalIndex[0]:decimalIndex[1]])
	}
	return buf.String()
}

// formatWithCommas is copied from ex 3.10
func formatWithCommas(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	start := n % 3
	buf := bytes.Buffer{}
	buf.WriteString(s[:start])
	if start > 0 {
		buf.WriteString(",")
	}
	for i := start; i < n-3; i += 3 {
		buf.WriteString(s[i : i+3])
		buf.WriteString(",")
	}
	buf.WriteString(s[n-3:])
	return buf.String()
}
