// Issues 는 깃헙 이슈를 테이블로 출력한다. 검색 텀과  매칭이 되는 녀석들로
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch4/_jeonghyunseok/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

/*
set GOPATH=d:\golang
go build -o issues.exe
issues.exe repo:golang/go is:open json decoder

*/
