package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "Hello World, \u4e16\u754c"
	data := []byte(s)

	//l := 0
	//fmt.Printf("l=%d, Len:=%d \n", l, len(data))

	for i := 0; i < len(data); {
		r, l := utf8.DecodeRune(data[i:])
		fmt.Printf("Char=%x  Len=%d, Counter:=%d \n", r, l, i)
		i += l
	}

	for i := 0; i < len(data); i++ {
		fmt.Printf("Char=%x, Counter:=%d \n", data[i:i+1], i)

	}

	data = reverse(data)

	//rune count - count all runes
	count := utf8.RuneCount(data)
	fmt.Printf("rune count = %d\n", count)

	//Decode last rune
	lr, lrl := utf8.DecodeLastRune(data)
	fmt.Printf("Last Rune = %c   Last rune len=%d\n", lr, lrl)

	fmt.Printf("Last Rune(Bytes) = %x   Start byte = %d\n", data[len(data)-lrl:], len(data)-lrl)

}

func reverse(s []byte) []byte {

	rs := s[:0] // Create a zero length slice of the original

	var buff []byte

	for i := len(s); i > 0; { // Start a loop on the len of orginal, handle i in loop

		lr, lrl := utf8.DecodeLastRune(s) //Find the last rune, lr, and its length, lrl

		n := utf8.EncodeRune(buff[:], lr)

		rs = append(rs, buff[:n]) // Append the last rune from s to rs

		for x := len(s) - lrl; x < len(s); x-- { //Loop to eliminate last rune
			s[x] = 0

		}

	}

	return rs

}
