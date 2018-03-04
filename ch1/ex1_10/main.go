// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
// Fetchall fetches URLs in parallel and reports their times and sizes.

// ex 1.10 find a website that produces a large amount of data. investigate
// caching by running fetchall twice in succession to see whether the reported
// time changes much. do you get the same content each time? modify fetchall to
// print it's output to a file so it can be examined.

// A: there is no significant difference in reported times over multiple runs,
// which makes sense because there is no local caching. the page views I tried
// must have already been cached on the web server.

// A: the content sometimes differs. this version of fetchall writes all the
// results into a single file fetchall.txt

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type result struct{ timing, content string }

const outFile = "fetchall.txt"

func main() {
	out, err := os.Create(outFile)
	defer out.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"fetchall: while opening %s, got %s", outFile, err)
	}
	start := time.Now()
	ch := make(chan result)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		// receive from channel ch
		res := <-ch;
		fmt.Println(res.timing)
		fmt.Fprintf(out, res.content)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- result) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- result{fmt.Sprint(err), ""}
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- result{fmt.Sprintf("while reading %s: %v", url, err), ""}
		return
	}
	secs := time.Since(start).Seconds()
	res := result{
		fmt.Sprintf("%.2fs  %7d  %s", secs, len(bytes), url),
		string(bytes),
	}
	ch <- res
}
