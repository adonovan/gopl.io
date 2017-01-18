// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// excercise 5-14
//I choose to use breadthfirst to look at the file on my Disk Drive
//It was necessary to write a nuw functions 'extract' to extract the folder
//names from each directory.
//To run you need to supply a starting directory path
//i.e.  ex5-14 c:/users/admin/googledrive/go/
//I curretnly have a problem with access privlidges.  If the program necounts a folder
//That requries extra privlidges, it will stop with an 'Access is denied' error

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)

	for len(worklist) > 0 {

		//fmt.Printf("\nWorklist: %v\n", worklist)

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
func crawl(folder string) []string {
	fmt.Printf("\nCrawl: %v\n", folder)
	list, err := extract(folder)
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
// func extreace takes a folder name and returns a slice with all the directory
//listed in the folder. It ignores file and symbolic links
//Is was necessary to add the whole directory path to the slice
func extract(folder string) ([]string, error) {
	folders, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var dirs []string

	for _, file := range folders {

		if file.IsDir() {

			fi, err := os.Lstat(folder + file.Name())
			if err != nil {
				log.Fatal(err)
				return nil, err
			}

			if (fi.Mode() & os.ModeSymlink) != os.ModeSymlink {
				//fmt.Printf("Extract: %v Mode:%v\n", file.Name(), file.Mode())
				fmt.Printf("Extract: %v\n", file.Name())
				dirs = append(dirs, folder+file.Name()+"/")

			}
		}

	}
	return dirs, nil
}
