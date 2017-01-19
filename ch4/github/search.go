// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))

	fmt.Printf("%s\n", strings.Join(terms, " ")) // DW: print the search strings

	fmt.Printf("%s\n", q) //DW: print the 'QueryEscaped' string. This seems to replace " " with "+"

	fmt.Printf("%s\n", IssuesURL+"?q="+q) //DW: print the url

	resp, err := http.Get(IssuesURL + "?q=" + q) //DW: 'resp' gets the complete response
	if err != nil {
		return nil, err
	}

	fmt.Println(reflect.TypeOf(resp))

	//!-
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	//   req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	//   if err != nil {
	//       return nil, err
	//   }
	//   req.Header.Set(
	//       "Accept", "application/vnd.github.v3.text-match+json")
	//   resp, err := http.DefaultClient.Do(req)
	//!+

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	//DW: In the following code:  'result' (Type IssuesSearchResult)
	// Call a json newDecoder, passing respon.Body, and the address of 'result'

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//!-
