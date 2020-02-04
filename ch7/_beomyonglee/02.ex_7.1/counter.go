package counter

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type LineCounter struct {
	lines int
}

func (c *LineCounter) Write(p []byte) (n int, err error) {
	for _, b := range p {
		if b == '\n' {
			c.lines++
		}
	}
	return len(p), nil
}

func (c *LineCounter) N() int {
	return c.lines
}

func (c *LineCounter) String() string {
	return fmt.Sprintf("%d", c.lines)
}

type WordCounter struct {
	words  int
	inWord bool
}

func leadingSpaces(p []byte) int {
	count := 0
	cur := 0
	for cur < len(p) {
		r, size := utf8.DecodeRune(p[cur:])
		if !unicode.IsSpace(r) {
			return count
		}
		cur += size
		count++
	}
	return count
}

func leadingNonSpaces(p []byte) int {
	count := 0
	cur := 0
	for cur < len(p) {
		r, size := utf8.DecodeRune(p[cur:])
		if unicode.IsSpace(r) {
			return count
		}
		cur += size
		count++
	}
	return count
}

func (c *WordCounter) Write(p []byte) (n int, err error) {
	cur := 0
	n = len(p)
	for {
		spaces := leadingSpaces(p[cur:])
		cur += spaces
		if spaces > 0 {
			c.inWord = false
		}
		if cur == len(p) {
			return
		}
		if !c.inWord {
			c.words++
		}
		c.inWord = true
		cur += leadingNonSpaces(p[cur:])
		if cur == len(p) {
			return
		}
	}
}

func (c *WordCounter) N() int {
	return c.words
}

func (c *WordCounter) String() string {
	return fmt.Sprintf("%d", c.words)
}
