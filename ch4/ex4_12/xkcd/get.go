/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package xkcd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Get fetches the comid with id and returns Comic a struct. If id is 0, fetches
// the most recent comic using the special URL https://xkcd.com/info.0.json
func Get(id int) (Comic, error) {
	var url string
	if id == 0 {
		url = fmt.Sprintf(APICurrentComic)
	} else {
		url = fmt.Sprintf(APIFormat, id)
	}
	resp, err := http.Get(url) // TODO: add timeouts to http client; dont use defaults
	if err != nil {
		return Comic{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		limitedReader := &io.LimitedReader{R: resp.Body, N: KiB}
		msg, _ := ioutil.ReadAll(limitedReader)
		err := fmt.Errorf(
			"comic get failed: status %s, msg %s",
			resp.Status,
			msg,
		)
		return Comic{}, err
	}
	limitedReader := &io.LimitedReader{R: resp.Body, N: MiB}
	data, err := ioutil.ReadAll(limitedReader)
	if err != nil {
		return Comic{}, err
	}
	comic := Comic{}
	err = json.Unmarshal(data, &comic)
	if err != nil {
		return Comic{}, err
	}
	return comic, nil
}
