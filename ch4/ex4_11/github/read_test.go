package github

import (
	"os"
	"testing"
)

func TestReadIssue(t *testing.T) {
	repo := Repo("guidorice/gopl.io")
	token := Token(os.Getenv(GithubEnvVar))
	issueId := IssueId(4)
	if token == "" {
		t.Fatal("missing Github token")
	}
	issue, err := ReadIssue(token, repo, issueId)
	if err != nil {
		t.Errorf("CreateIssue: %v", err)
	}
	t.Logf("%v", issue)
}
