package github

import (
	"os"
	"testing"
)

func TestCreateIssue(t *testing.T) {
	repo := Repo("guidorice/gopl.io")
	token := Token(os.Getenv(GithubEnvVar))
	if token == "" {
		t.Fatal("missing Github token")
	}
	issue := IssueCreate{
		Title: "test issue",
		Body:  "lorem ipsum",
	}
	_, err := CreateIssue(token, repo, issue)
	if err != nil {
		t.Errorf("CreateIssue: %v", err)
	}
}
