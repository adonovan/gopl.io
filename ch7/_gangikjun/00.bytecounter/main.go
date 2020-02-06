package main

import (
	"fmt"
)

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0
	var name = "Dooly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

// ByteCounter ByteCounter
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
