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
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
