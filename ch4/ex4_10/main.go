/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

/* Issues prints a table of GitHub issues matching the search terms ; ex 4.10:
 * modify issues to report the results in age categories, say less than a month
 * old, less than a year old, and more than a year old. Note issues will be
 * categorized as either PastMonth or PastYear, but not both.

$ go run main.go repo:golang/go is:open json encoder

20 issues:

**Created in past month:**

**Created in past year:**
#22967 2017-12-01 21:27:06 +0000 UTC,    larhun encoding/json: bad encoding of field with MarshalJSON m
#25153 2018-04-29 03:44:33 +0000 UTC, dominicba encoding/json: include property name for marshal errors
#22480 2017-10-28 20:20:06 +0000 UTC,     blixt proposal: encoding/json: add omitnil option
#26264 2018-07-07 19:13:53 +0000 UTC,     mvdan bytes: WriteString, WriteByte and Write vary in perform
#25426 2018-05-16 15:25:37 +0000 UTC, josharian cmd/compile: revisit statement boundaries CL peformance

**Older:**
#7872  2014-04-26 17:47:25 +0000 UTC, extempora encoding/json: Encoder internally buffers full output
#12001 2015-08-03 19:14:17 +0000 UTC, lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
#14140 2016-01-28 22:40:17 +0000 UTC,    rgooch encoding/gob, encoding/json: add streaming interface
#5901  2013-07-17 16:39:15 +0000 UTC,       rsc encoding/json: allow override type marshaling
#5819  2013-06-30 19:00:22 +0000 UTC, gopherbot encoding/gob: encoder should ignore embedded structs wi
#14804 2016-03-13 15:06:10 +0000 UTC, cyberphon encoding/json: marshaller does not provide Base64Url su
#17609 2016-10-26 16:07:27 +0000 UTC, nathanjsw encoding/json: ambiguous fields are marshalled
#11046 2015-06-03 19:25:08 +0000 UTC,     kurin encoding/json: Decoder internally buffers full input
#10769 2015-05-10 19:13:32 +0000 UTC,       bep encoding/json: detect circular data structures when enc
#11939 2015-07-30 14:14:34 +0000 UTC,   joeshaw proposal: encoding/json, encoding/xml: support zero val
#20754 2017-06-22 14:19:31 +0000 UTC,       rsc encoding/xml: unmarshal only processes first XML elemen
#20206 2017-05-02 09:15:12 +0000 UTC, markdryan encoding/base64: encoding is slow
#15808 2016-05-24 02:13:10 +0000 UTC, randall77 cmd/compile: loads/constants not lifted out of loop
#4146  2012-09-25 00:47:27 +0000 UTC,  bradfitz reflect: MakeInterface
#17244 2016-09-27 00:41:32 +0000 UTC,       adg proposal: decide policy for sub-repositories

*/

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

type Category int

const (
	PastMonth Category = iota
	PastYear
	Older
)

var issuesByAge = make(map[Category][]*github.Issue)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		dur := time.Since(item.CreatedAt)
		days := int(dur.Hours() / 24)
		if days <= 31 {
			issuesByAge[PastMonth] = append(issuesByAge[PastMonth], item)
		} else if days <= 365 {
			issuesByAge[PastYear] = append(issuesByAge[PastYear], item)
		} else {
			issuesByAge[Older] = append(issuesByAge[Older], item)
		}
	}

	fmt.Printf("\n**Created in past month:**\n")
	printIssues(PastMonth)
	fmt.Printf("\n**Created in past year:**\n")
	printIssues(PastYear)
	fmt.Printf("\n**Older:**\n")
	printIssues(Older)
}

func printIssues(c Category) {
	for _, item := range issuesByAge[c] {
		fmt.Printf("#%-5d %v, %9.9s %.55s\n",
			item.Number, item.CreatedAt, item.User.Login, item.Title)
	}
}
