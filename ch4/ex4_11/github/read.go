/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package github

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// ReadIssue reads an issue id from github api
func ReadIssue(token Token, repo Repo, id IssueId) (Issue, error) {
	url := APIURL + "repos/" + string(repo) + "/issues/" + string(id)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", string(url), nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+string(token))
	resp, err := client.Do(req)
	if err != nil {
		return Issue{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		limitedReader := &io.LimitedReader{R: resp.Body, N: KiB}
		msg, _ := ioutil.ReadAll(limitedReader)
		err := fmt.Errorf(
			"create read failed: status %s, msg %s",
			resp.Status,
			msg,
		)
		return Issue{}, err
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
