package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Printf("index :%d, value :%s\n", i, v)
	}
}
