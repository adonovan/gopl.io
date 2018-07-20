/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// ex 4.6 Write an in-place function that squashes each run of adjacent unicode
// spaces (see unicode.IsSpace) in a utf8 encoded []byte slice into a single
// ascii space.

package squash

import (
	"bytes"
	"unicode"
)

type sliceOfBytes []byte

// SquashSpaces squashes each run of adjacent unicode spaces in the receiver.
// Each run of unicode spaces is converted to a single ascii space.
func (r *sliceOfBytes) SquashSpaces() {
	// convert to runes in preparation for calling unicode.IsSpace
	buf := bytes.Buffer{}
	runes := bytes.Runes(([]byte)(*r))
	squashing := false
	for _, rn := range runes {
		if unicode.IsSpace(rn) {
			if !squashing {
				buf.WriteString(" ")
				squashing = true
			}
			continue
		}
		buf.WriteRune(rn)
		squashing = false
	}
	*r = buf.Bytes()
}
