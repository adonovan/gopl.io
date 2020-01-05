package main

import (
	"./github"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1:]
	// var result IssueSearchResult
	fmt.Println("Main Program started")
	result, err := github.SearchIssue(input)

	if err != nil {
		fmt.Printf("%v", result)
	}
}


