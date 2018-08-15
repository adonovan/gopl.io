/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateIssue(
	token GithubToken,
	repo string,
	issue IssueTemplate,
) (string, error) {
	url := APIURL + "repos/" + repo + "/issues"
	b, err := json.Marshal(issue)
	if err != nil {
		return "", err
	}
	body := bytes.NewBuffer(b)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+string(token))
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		msg, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("create issue failed: status %s, msg %s", resp.Status, msg)
	}
	return "", nil
}
