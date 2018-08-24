// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"time"
)

const APIURL = "https://api.github.com/"
const GithubEnvVar = "GITHUB_TOKEN"

var IssuesSearchURL = APIURL + "search/issues"

// Token is a string type for github authorization token.
type Token string

// Repo is a string type for github repository path.
type Repo string

// IssueId is a string type for convenience (although issue ids are sequential integers)
type IssueId string

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// IssueCreate is a struct for json serialization from github api. It
// has a slightly different layout than the full Issue struct returned by the api.
type IssueCreate struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	State     string   `json:"string,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}
