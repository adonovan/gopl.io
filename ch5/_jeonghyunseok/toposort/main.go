// Toposort 는 DAG 를 토폴로지 순서로 출력해준다.
// Directed Acycle Graph

package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structrues"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) // 이걸 안해놓으면 재귀적으로 못쓴다.

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	fmt.Println("sorted:", keys)
	visitAll(keys)
	return order
}

/*
1) main 에서 topoSort(prereqs) 의 리턴값을 loop 돌면서 출력한다.
2) topoSort 는
	- 우선 map 의 key 값들을 다 슬라이스에 넣어서 소팅한 다음에
	- 그걸 visitAll() 해버린다.
3) visitAll() 은
	- 일단 자신을 추가하고, 자신의 자식들을 방문한다.
	- 자식들 방문은 정렬되지 않은것도 같다.근데 자식들도 부모에 이미 있는 것들이라 상관없을듯

go run main.go



*/

// 흠 근데 왜 순서가 algorithm 이 먼저가 아닐까 ?
// visitAll() 해서 계속 파고 들어가다가 더 못들어가면 시작하기 때문

// sorted: [algorithms calculus compilers data structures databases discrete math formal languages networks operating systems programming languages]
// 1:      data structrues
// 2:      algorithms
// 3:      linear algebra
// 4:      calculus
// 5:      intro to programming
// 6:      discrete math
// 7:      data structures
// 8:      formal languages
// 9:      computer organization
// 10:     compilers
// 11:     databases
// 12:     operating systems
// 13:     networks
// 14:     programming languages
