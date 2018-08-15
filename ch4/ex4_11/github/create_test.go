package github

import (
	"os"
	"testing"
)

func TestCreateIssue(t *testing.T) {
	repo := "guidorice/gopl.io"
	token := GithubToken(os.Getenv(GithubEnvVar))
	if token == "" {
		t.Fatal("missing Github token")
	}
	issue := IssueTemplate{
		Title: "test issue",
		Body:  "lorem ipsum",
	}
	_, err := CreateIssue(token, repo, issue)
	if err != nil {
		t.Errorf("CreateIssue: %v", err)
	}
}
