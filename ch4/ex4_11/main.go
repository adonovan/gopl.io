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

// runtime options (after flag parse)
type options struct {
	action  action
	repo    string
	issue   string
	message string // can also be search terms, space separated
	token   github.GithubToken
}

// container for flag parsing
type flags struct {
	search bool
	create bool
	read   bool
	update bool
	close  bool
}

var opts options
var f flags

// package init
func init() {
	flag.BoolVar(&f.create, "create", true, "create a github issue")
	flag.StringVar(&opts.repo, "repo", "", "repository name")
}

// parseAction converts boolean flags into more readable action type.
func parseAction() {
	if f.create {
		opts.action = create
	}
}

func main() {
	flag.Parse()
	parseAction()
	opts.token = github.GithubToken(os.Getenv(github.GithubEnvVar))
	if opts.token == "" {
		log.Fatal("error: missing github token in env " + github.GithubEnvVar)
	}
	if opts.repo == "" {
		log.Fatal("error: missing repository name")
	}
	switch opts.action {
	case create:
		issue := github.IssueTemplate{
			Title: "test api create",
			Body:  "lorem ipsum",
		}
		github.CreateIssue(opts.token, opts.repo, issue)
	}
}
