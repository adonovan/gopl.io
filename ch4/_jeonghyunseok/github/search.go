// github is
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println("q:\n", q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	/*
		req, err := http.NewReuquest("GET", IssuesURL+"?q="+q, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set(
			"Accept", "application/vnd.github.v3.text-matech+json")
			resp, err := http.DefaultClient.Do(req)
		)
	*/

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil

}

// defer is much better
// issues example code will use this package
