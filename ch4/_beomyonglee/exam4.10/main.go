// 연습문제 4.10 issues를 수정해 한달 미만, 1년 미만, 1년 이상과 같이
// 특정 수명 범주에 속하는 결과를 출력하라.
package main

import (
	"fmt"
	"gopl.io/ch4/_beomyonglee/github"
	"log"
	"os"
	"time"
)

/*
특정 수명 범주를 나타내는 검색 qualifier가 있는가?
*/
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	format := "#%-5d %9.9s %.55s\n"
	now := time.Now()

	pastDay := make([]*github.Issue, 0)
	pastMonth := make([]*github.Issue, 0)
	pastYear := make([]*github.Issue, 0)

	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(day):
			pastDay = append(pastDay, item)
		case item.CreatedAt.After(month) && item.CreatedAt.Before(day):
			pastMonth = append(pastMonth, item)
		case item.CreatedAt.After(year) && item.CreatedAt.Before(month):
			pastYear = append(pastYear, item)
		}
	}

	if len(pastDay) > 0 {
		fmt.Printf("\nPast day:\n")
		for _, item := range pastDay {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
	if len(pastMonth) > 0 {
		fmt.Printf("\nPast month:\n")
		for _, item := range pastMonth {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
	if len(pastYear) > 0 {
		fmt.Printf("\nPast year:\n")
		for _, item := range pastYear {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
}

/*
$ .go build main.go
$ ./main repo:golang/go is:open json decoder

Past month:
#36225     dsnet encoding/json: the Decoder.Decode API lends itself to m
#36353     dsnet proposal: encoding/gob: allow override type marshaling

Past year:
#33416   bserdar encoding/json: This CL adds Decoder.InternKeys
#34647 babolivie encoding/json: fix byte counter increments when using d
#34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report
#32779       rsc proposal: encoding/json: memoize strings during decode?
#34564  mdempsky go/internal/gcimporter: single source of truth for deco
#31789  mgritter encoding/json: provide a way to limit recursion depth
#31701    lr1980 encoding/json: second decode after error impossible
#30301     zelch encoding/xml: option to treat unknown fields as an erro

./main repo:golang/go "created:>2015-01-01"
./main repo:golang/go created:2015-01-01
./main repo:golang/go created:2015-01-01 created:2015-01-02
*/
