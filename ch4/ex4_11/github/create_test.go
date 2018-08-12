package github

import "testing"

func TestCreateIssue(t *testing.T) {
	repo := "guidorice/gopl.io"
	issue := &IssueCreate{
		Title: "test issue",
		Body:  "lorem ipsum",
	}
	_, err := CreateIssue(repo, issue)
	if err != nil {
		t.Errorf("CreateIssue: %v", err)
	}
}
