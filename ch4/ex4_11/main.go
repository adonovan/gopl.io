/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// ex 4.11: Build a tool that lets users create, read, update, and close github
// issues from the command line, invoking their preferred text editor when sub-
// stantial text input is required.
package main

// TODO: add update feature w/ editor

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

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

// tmp container for flag parsing
type flags struct {
	search bool
	create bool
	read   int
	update int
	close  int
	repo   string
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
	flag.BoolVar(&f.create, "c", false, "create a github issue (shorthand)")
	flag.IntVar(&f.close, "close", -1, "close a github issue")
	flag.IntVar(&f.read, "read", -1, "read a github issue")
	flag.IntVar(&f.read, "r", -1, "read a github issue (shorthand)")
	flag.BoolVar(&f.search, "search", false, "search issues by terms")
	flag.BoolVar(&f.search, "s", false, "search issues by terms (shorthand)")
	flag.StringVar(&f.repo, "repo", "", "repository name")
	flag.Usage = Usage
}

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	// TODO: add detailed usage w/ examples
	flag.PrintDefaults()
}

// checkFlags converts boolean, int, etc flags into more readable types.
func checkFlags() {
	opts.repo = github.Repo(f.repo)
	if f.create {
		opts.action = create
	} else if f.read > 0 {
		opts.action = read
		opts.issue = github.IssueId(strconv.Itoa(f.read))
	} else if f.close > 0 {
		opts.action = close
		opts.issue = github.IssueId(strconv.Itoa(f.close))
		if opts.issue == "0" {
			flag.Usage()
			os.Exit(2)
		}
	} else if f.search {
		opts.action = search
		opts.message = strings.Join(flag.Args(), " ")
	}
}

const IssuesSearchTemplate = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var IssuesSearchReport = template.Must(template.New("searchResult").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(IssuesSearchTemplate))

const IssueTemplate = `issue#:  {{.Number}} {{.HTMLURL}}
title:   {{.Title}}
state:   {{.State}}
age:    {{.CreatedAt | daysAgo}} days
creator: {{.User.Login}} {{.User.HTMLURL}}

{{.Body}}
`

var IssueReport = template.Must(template.New("issue").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(IssueTemplate))

func main() {
	flag.Parse()
	checkFlags()
	opts.token = github.Token(os.Getenv(github.GithubEnvVar))
	if opts.token == "" || opts.repo == "" {
		flag.Usage()
		os.Exit(2)
	}
	switch opts.action {
	case create:
		// TODO: get title, body from flag
		// TODO: open editor option
		tmpl := github.IssueCreate{
			Title: "test api create",
			Body:  "lorem ipsum",
		}
		_, err := github.CreateIssue(opts.token, opts.repo, tmpl)
		if err != nil {
			log.Fatal(err)
		}
	case read:
		issue, err := github.ReadIssue(opts.token, opts.repo, opts.issue)
		if err != nil {
			log.Fatal(err)
		}
		IssueReport.Execute(os.Stdout, issue)
	case search:
		// prefix the search terms with a single repo
		terms := "repo:" + string(opts.repo) + " " + opts.message
		results, err := github.SearchIssues(terms)
		if err != nil {
			log.Fatal(err)
		}
		IssuesSearchReport.Execute(os.Stdout, results)
	case close:
		issue, err := github.CloseIssue(opts.token, opts.repo, opts.issue)
		if err != nil {
			log.Fatal(err)
		}
		IssueReport.Execute(os.Stdout, issue)
	}
}
