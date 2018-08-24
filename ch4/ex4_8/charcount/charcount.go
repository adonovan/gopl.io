/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package charcount

import "unicode"

type Counts struct {
	characters  int
	control     int
	digit       int
	graphic     int
	letter      int
	lowerCase   int
	mark        int
	number      int
	punctuation int
	space       int
	symbol      int
	upperCase   int
}

// Charcount counts the number of unicode characters and counts of 11 unicode
// character categories (upper, lower, symbol, etc).
func Charcount(s string) Counts {
	c := Counts{}
	runes := []rune(s)
	for _, r := range runes {
		c.characters++
		if unicode.IsControl(r) {
			c.control++
		}
		if unicode.IsDigit(r) {
			c.digit++
		}
		if unicode.IsGraphic(r) {
			c.graphic++
		}
		if unicode.IsLetter(r) {
			c.letter++
		}
		if unicode.IsLower(r) {
			c.lowerCase++
		}
		if unicode.IsMark(r) {
			c.mark++
		}
		if unicode.IsNumber(r) {
			c.number++
		}
		if unicode.IsPunct(r) {
			c.punctuation++
		}
		if unicode.IsSpace(r) {
			c.space++
		}
		if unicode.IsSymbol(r) {
			c.symbol++
		}
		if unicode.IsUpper(r) {
			c.upperCase++
		}
	}
	return c
}
