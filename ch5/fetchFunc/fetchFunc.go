// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 148.

// Fetch saves the contents of a URL into a local file.
package fetchFunc

import (
	//"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

//!+
// Fetch downloads the URL and returns the
// name and length of the local file.
func Fetch(url string) (filename string, n int64, err error) {
	//	resp, err := http.Head(url)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	//fmt.Printf("Request: %v\n\n", resp.Request)
	//fmt.Printf("Request.URL: %v\n\n", resp.Request.URL)
	//fmt.Printf("Request.URL.Path: %v\n\n", resp.Request.URL.Path)
	//fmt.Printf("path.Base(Request.URL.Path): %v\n\n", path.Base(resp.Request.URL.Path))

	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}


