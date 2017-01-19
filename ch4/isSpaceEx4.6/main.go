package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {

	s := "Hello   World"
	data := []byte(s)

	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(data); i++ {
		r, len := utf8.DecodeRune(data[i:]) //Decodes the first rune found in a []byte
		if len == 0 {                       // If the lengthof the rune is 0 quit
			break // and if you don't use len, you'll get and error
		}

		r2, len2 := utf8.DecodeRune(data[i+1:]) //Decode the next rune
		if len2 == 0 {
			break
		}

		result := unicode.IsSpace(r) //Check to see if the runes are spaces
		result2 := unicode.IsSpace(r2)

		// if the current char and next chars are spaces set the next char to 0
		if result && result2 {
			data[i] = 0
		}

	}

	for i := 0; i < len(data); i++ {
		fmt.Printf("Data[%d] = %c\n", i, data[i])
	}

	// Rmmove the duplicates
	out := data[:0] // Create a zero length slice of the original

	for _, s := range data { // lose the nils and build new slice 'out'
		if s != 0 {
			out = append(out, s)
		}

	}

	for i := 0; i < len(data); i++ {
		fmt.Printf("Data[%d] = %c\n", i, data[i])
	}
}
