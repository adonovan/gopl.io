/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package github

import (
	"os"
	"strings"
	"testing"
)

func TestSearchIssues(t *testing.T) {
	terms := []string{
		"repo:golang/go",
		"is:open",
		"json",
		"decoder",
	}
	token := Token(os.Getenv(GithubEnvVar))
	_, err := SearchIssues(token, strings.Join(terms, " "))
	if err != nil {
		t.Errorf("SearchIssues: %v", err)
	}
}
