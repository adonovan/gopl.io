// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

// !+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	issues := make(map[string][]*github.Issue)
	issues["less than a month old"] = make([]*github.Issue, 0)
	issues["less than a year old"] = make([]*github.Issue, 0)
	issues["older than a year"] = make([]*github.Issue, 0)
	now := time.Now()
	month := now.Add(-30 * 24 * time.Hour)
	year := now.Add(-365 * 24 * time.Hour)
	for _, item := range result.Items {
		if item.CreatedAt.After(month) {
			issues["less than a month old"] = append(issues["less than a month old"], item)
		} else if item.CreatedAt.Before(month) && item.CreatedAt.After(year) {
			issues["less than a year old"] = append(issues["less than a year old"], item)
		} else {
			issues["older than a year"] = append(issues["older than a year"], item)
		}
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Print("\n[Issues less than a month old]\n")
	for _, item := range issues["less than a month old"] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Print("\n[Issues less than a year old]\n")
	for _, item := range issues["less than a year old"] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Print("\n[Issues older than a year]\n")
	for _, item := range issues["older than a year"] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
