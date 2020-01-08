// Issues는 검색어와 일치하는 GitHub 이슈의 테이블을 출력한다.
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch4/_beomyonglee/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("|#%-5d|%9.9s|%.55s|\n", //-5: -는 좌측정렬, 5는 너비| 9.9: 점 앞 9는 너비, 점 뒤 9는 글자수
			item.Number, item.User.Login, item.Title)
	}
}

/*
go build main.go
./main repo:golang/go is:open json decoder
*/
