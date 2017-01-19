// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

//Modified by Douglas Will
//

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/dugwill/gopl.io/ch5/fetchFunc"
	"github.com/dugwill/gopl.io/ch5/links"
)

var origHost string

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {

	u, err := url.Parse(worklist[0])
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Worklist: %v\n", u.Host)
	origHost = u.Host

	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(theUrl string) []string {

	u, err := url.Parse(theUrl)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(theUrl)
	if u.Host == origHost {
		fmt.Println(theUrl)

		local, n, err := fetchFunc.Fetch(theUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", theUrl, err)
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", theUrl, local, n)

	}

	list, err := links.Extract(theUrl)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
