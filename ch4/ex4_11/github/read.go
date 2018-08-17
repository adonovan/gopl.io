/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package github

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// TODO: unmarhall the issue, and change return type to struct.

// ReadIssue reads an issue using github api and the GithubToken, repo
func ReadIssue(token Token, repo Repo, issueNumber IssueId) ([]byte, error) {
	url := APIURL + "repos/" + string(repo) + "/issues/" + strconv.Itoa(int(issueNumber))
	log.Printf("1: %v  2:%v", issueNumber, string(issueNumber))
	client := &http.Client{}
	req, _ := http.NewRequest("GET", string(url), nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+string(token))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		limitedReader := &io.LimitedReader{R: resp.Body, N: KiB}
		msg, _ := ioutil.ReadAll(limitedReader)
		return nil, fmt.Errorf("create read failed: status %s, msg %s", resp.Status, msg)
	}
	limitedReader := &io.LimitedReader{R: resp.Body, N: MiB}
	msg, err := ioutil.ReadAll(limitedReader)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
