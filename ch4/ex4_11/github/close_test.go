package github

import (
	"os"
	"testing"
)

func TestCloseIssue(t *testing.T) {
	repo := Repo("guidorice/gopl.io")
	token := Token(os.Getenv(GithubEnvVar))
	issueId := IssueId("16")
	if token == "" {
		t.Fatal("missing Github token")
	}
	_, err := CloseIssue(token, repo, issueId)
	if err != nil {
		t.Errorf("CloseIssue: %v", err)
	}
}
