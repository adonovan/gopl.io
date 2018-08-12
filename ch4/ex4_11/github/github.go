// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import "time"

const APIURL = "https://api.github.com/"

var IssuesSearchURL = APIURL + "search/issues"

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

// IssueCreate is the payload for creating an issue by POST /repos/:owner/:repo/issues
type IssueCreate struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Assignees []string `json:"assignees,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}
