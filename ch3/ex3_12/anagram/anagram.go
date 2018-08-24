/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package anagram

// Anagram returns true or false whether two strings, a and b, are anagrams.
func Anagram(a string, b string) bool {
	if a == b {
		return false
	}
	seen := map[string]bool{} // set of strings
	for _, c := range a {
		s := string(c)
		seen[s] = true
	}
	for _, c := range b {
		s := string(c)
		if !seen[s] {
			return false
		}
	}
	return true
}
