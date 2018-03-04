// unicodetest: confirm that unicode character gets automatically parsed and
//  displayed correctly when using byte slices.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Go is a general-purpose language designed with " +
		"systems programming in mind. 🥑\n")
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}
