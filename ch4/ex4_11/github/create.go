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

const token = "xxx" // FIXME get token from cli

func CreateIssue(repo string, issue *IssueCreate) (string, error) {
	url := APIURL + "repos/" + repo + "/issues"
	fmt.Println(url)
	b, err := json.Marshal(issue)
	if err != nil {
		return "", err
	}
	fmt.Println(string(b))
	body := bytes.NewBuffer(b)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+token)
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
