package main

import (
	"fmt"


	"./templates"
)

func main() {
	fmt.Printf("%s\n", templates.Hello("Foo"))
	fmt.Printf("%s\n", templates.Hello("Bar"))

}
