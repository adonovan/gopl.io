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

// SearchIssues queries the GitHub issue tracker. note: searches across all
// repos, so add add terms like repo:golang/go is:open json decoder to narrow scope.
func SearchIssues(token Token, terms string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(terms)
	u := IssuesSearchURL + "?q=" + q
	client := &http.Client{}
	req, _ := http.NewRequest("GET", string(u), nil)
	// allow authenticated searches, but don't require token
	if string(token) != "" {
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "token "+string(token))
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("SearchIssues: search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
