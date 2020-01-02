// Charcount는 유니코드 문자의 카운트를 계산한다
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // 유니코드 문자 카운트
	var utflen [utf8.UTFMax + 1]int // UTF-8 인코딩 길이 카운트
	invalid := 0                    // 잘못된 UTF-8 문자 카운트

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // rune, nbytes, error 반환
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("charcount: %v\n", err)
		}
		if !unicode.IsLetter(r) || r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
