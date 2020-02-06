package main

import (
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

// 동시 크롤러에 깊이 제한 추가

var tokens = make(chan struct{}, 20)
var maxdepth = 3

type work struct {
	url   string
	depth int
}

func crawl(w work) []work {
	fmt.Println("depth:", w.depth, ", url:", w.url)

	if w.depth >= maxdepth {
		return nil
	}

	tokens <- struct{}{}
	list, err := links.Extract(w.url)
	<-tokens

	if err != nil {
		log.Print(err)
	}

	newwork := []work{}
	for _, link := range list {
		newwork = append(newwork, work{link, w.depth + 1})
	}

	return newwork
}

func main() {
	worklist := make(chan []work)
	var n int

	n++
	go func() {
		w := []work{}
		w = append(w, work{"http://gopl.io", 1})
		worklist <- w
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				n++
				go func(link work) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
