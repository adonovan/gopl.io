// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	
	file, err := os.Open("github.com/InputFiles/manifest.mpd") // For read access.
	if err != nil {
		//log.Fatal(err)
		fmt.Printf("Error opening file\n")
}


	//dec := xml.NewDecoder(os.Stdin)

	dec := xml.NewDecoder(file)

	var stack []string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			fmt.Printf("StartElement: %v\n",tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
			fmt.Printf("EndElement\n")
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		case xml.Comment:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		case xml.Decoder:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			fmt.Printf("Returning True...................\n")
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}

	fmt.Printf("Returning False\n")
	return false
}

//!-
