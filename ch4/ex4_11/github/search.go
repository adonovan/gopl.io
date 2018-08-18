/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

//!+

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// TODO: add auth token so private repos can be searched as well as public.

// SearchIssues queries the GitHub issue tracker. note: searches across all
// repos, so add add terms like repo:golang/go is:open json decoder to narrow scope.
func SearchIssues(terms string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(terms)
	resp, err := http.Get(IssuesSearchURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
