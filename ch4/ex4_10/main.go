/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// Issues prints a table of GitHub issues matching the search terms ; ex 4.10:
// modify issues to report the results in age categories, say less than a month
// old, less than a year old, and more than a year old. Note issues will be
// categorized as either PastMonth or PastYear, but not both.

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

type Category int

const (
	PastMonth Category = iota
	PastYear
	Older
)

var issuesByAge = make(map[Category][]*github.Issue)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n\n", result.TotalCount)

	for _, item := range result.Items {
		dur := time.Since(item.CreatedAt)
		days := int(dur.Hours() / 24)
		if days <= 31 {
			issuesByAge[PastMonth] = append(issuesByAge[PastMonth], item)
		} else if days <= 365 {
			issuesByAge[PastYear] = append(issuesByAge[PastYear], item)
		} else {
			issuesByAge[Older] = append(issuesByAge[Older], item)
		}
	}

	fmt.Printf("\n**Created in past month:**\n")
	printIssues(PastMonth)
	fmt.Printf("\n**Created in past year:**\n")
	printIssues(PastYear)
	fmt.Printf("\n**Older:**\n")
	printIssues(Older)
}

func printIssues(c Category) {
	for _, item := range issuesByAge[c] {
		fmt.Printf("#%-5d %v, %9.9s %.55s\n",
			item.Number, item.CreatedAt, item.User.Login, item.Title)
	}
}
