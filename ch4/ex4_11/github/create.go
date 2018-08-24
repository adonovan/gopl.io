/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// CreateIssue creates an issue using github api and the token, repo, and
// template.
func CreateIssue(token Token, repo Repo, issueSrc IssueCreate) (Issue, error) {
	url := APIURL + "repos/" + repo + "/issues"
	b, err := json.Marshal(issueSrc)
	if err != nil {
		return Issue{}, err
	}
	body := bytes.NewBuffer(b)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", string(url), body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+string(token))
	resp, err := client.Do(req)
	if err != nil {
		return Issue{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		limitedReader := &io.LimitedReader{R: resp.Body, N: KiB}
		msg, _ := ioutil.ReadAll(limitedReader)
		return Issue{}, fmt.Errorf("CreateIssue: failed with status %s, msg %s", resp.Status, msg)
	}
	limitedReader := &io.LimitedReader{R: resp.Body, N: MiB}
	data, err := ioutil.ReadAll(limitedReader)
	if err != nil {
		return Issue{}, err
	}
	issue := Issue{}
	err = json.Unmarshal(data, &issue)
	if err != nil {
		return Issue{}, err
	}
	return issue, nil
}
