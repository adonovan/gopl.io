/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// ex 4.11: Build a tool that lets users create, read, update, and close github
// issues from the command line, invoking their preferred text editor when sub-
// stantial text input is required.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/guidorice/gopl.io/ch4/ex4_11/github"
)

type action int

const (
	undefined action = iota
	search
	create
	read
	update
	close
)

// container for flag parsing
type flags struct {
	search bool
	create bool
	read   bool
	update bool
	close  bool
	repo   string
	issue  int
}

// runtime options (after flag parse)
type options struct {
	action  action
	repo    github.Repo
	issue   github.IssueId
	message string // can also be search terms, space separated
	token   github.Token
}

var opts options
var f flags

// package init
func init() {
	flag.BoolVar(&f.create, "create", false, "create a github issue")
	flag.BoolVar(&f.read, "read", false, "read a github issue")
	flag.StringVar(&f.repo, "repo", "", "repository name")
	flag.IntVar(&f.issue, "issue", 0, "issue number")
}

// parseAction converts boolean flags into more readable action type.
func parseAction() {
	opts.repo = github.Repo(f.repo)
	opts.issue = github.IssueId(f.issue)
	if f.create {
		opts.action = create
	} else if f.read {
		opts.action = read
	}
}

func main() {
	flag.Parse()
	parseAction()
	opts.token = github.Token(os.Getenv(github.GithubEnvVar))
	if opts.token == "" {
		log.Fatal("error: missing github token in env " + github.GithubEnvVar)
	}
	if opts.repo == "" {
		log.Fatal("error: missing repository name")
	}
	switch opts.action {
	case create:
		template := github.IssueCreateTemplate{
			Title: "test api create",
			Body:  "lorem ipsum",
		}
		_, err := github.CreateIssue(opts.token, opts.repo, template)
		if err != nil {
			log.Fatal(err)
		}
	case read:
		issue, err := github.ReadIssue(opts.token, opts.repo, opts.issue)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v", issue)
	}
}
