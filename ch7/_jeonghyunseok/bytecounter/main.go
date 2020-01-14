// Bytecounter 는 바이트를 세는 io.Writer 를 구현한다.
package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	c = 0
	var koreanName = "정현석"
	c.Write([]byte(koreanName))
	fmt.Println(c)
}

// go run main.go
