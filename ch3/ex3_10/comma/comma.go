/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package comma

import "bytes"

// comma inserts commas in a non-negative decimal integer string.
func Comma(s string) string {
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
	for i := start; i < n - 3; i += 3 {
		buf.WriteString(s[i:i+3])
		buf.WriteString(",")
	}
	buf.WriteString(s[n-3:])
	return buf.String()
}
