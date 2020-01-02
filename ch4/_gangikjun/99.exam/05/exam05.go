package main

import "fmt"

func main() {
	s := []string{"a", "a", "b", "b"}

	fmt.Println(unique(s))
}

func nearStringModify(s []string) []string {
	var result []string
	for i := range s {
		if i+1 < len(s) {
			if s[i] == s[i+1] {
				s = remove(s, i)
				result = s
			}
		}
	}
	return result
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func unique(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s == strings[i] {
			continue
		}
		i++
		strings[i] = s
	}
	return strings[:i+1]
}
