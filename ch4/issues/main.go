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

	"github.com/dugwill/MyProjects/timeCalc"
	"github.com/dugwill/gopl.io/ch4/github"
)

const ThirtyDays time.Duration = 720 * time.Hour //Set var month to 720 hours

//!+
func main() {

	result, err := github.SearchIssues(os.Args[1:]) // DW: 'result' is a struct of type IssuesSearchResult
	// SearchIssues returns the address of the struct
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	t0 := time.Now()
	var yearAgo time.Time = t0.Add(-timeCalc.OneYear)
	yearStart := time.Date(2016, 1, 1, 00, 00, 00, 00, time.UTC)
	fmt.Println(yearAgo.Format("One Year ago was: Jan 2 15:04:05 2006 MST"))

	fmt.Println()
	fmt.Println("Items created this year")

	for _, item := range result.Items { // DW: 'result.Items' is a slice containing the pointers to issues

		//Dispaly items created within the last year
		if item.CreatedAt.After(yearStart) {
			fmt.Printf("#%-5d %9.9s %v %.55s\n", //DW: Add %v to display the time
				item.Number, item.User.Login, item.CreatedAt, item.Title) //DW: added item.CreatedAt to display time
		}

	}
	fmt.Println()
	fmt.Println("Items less than one year old")

	for _, item := range result.Items { // DW: 'result.Items' is a slice containing the pointers to issues

		//Dispaly items created within the last year
		if item.CreatedAt.After(yearAgo) && item.CreatedAt.Before(yearStart) {
			fmt.Printf("#%-5d %9.9s %v %.55s\n", //DW: Add %v to display the time
				item.Number, item.User.Login, item.CreatedAt, item.Title) //DW: added item.CreatedAt to display time
		}

	}

	fmt.Println()
	fmt.Println("Items greater than one year old")

	for _, item := range result.Items { // DW: 'result.Items' is a slice containing the pointers to issues

		//Dispaly items created within the last year
		if item.CreatedAt.Before(yearAgo) {
			fmt.Printf("#%-5d %9.9s %v %.55s\n", //DW: Add %v to display the time
				item.Number, item.User.Login, item.CreatedAt, item.Title) //DW: added item.CreatedAt to display time
		}

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
