package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println("index: ", index, "argument: ", arg)

	}

}
