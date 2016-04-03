//Excercise 4.5
// Removes adjacent duplicates in a slice

package main

import "fmt"

func main() {

	data := []string{"one", "one", "two", "three", "four", "four", "one"}

	data = rmdup(data)

	fmt.Printf("%q\n", data)
}

func rmdup(strings []string) []string {

	out := strings[:0] // Create a zero length slice of the original

	for i := 0; i < len(strings)-1; i++ { // Create a loop the length of s
		if strings[i+1] == strings[i] { // if there are dups set them to nil
			strings[i+1] = ""
			fmt.Printf("%q\n", strings)
		}
	}

	for _, s := range strings { // lose the nils and build new slice 'out'
		if s != "" {
			out = append(out, s)
		}

	}
	return out // return slice 'out'
}
