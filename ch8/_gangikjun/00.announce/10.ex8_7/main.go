package main

import "10.ex8_7/crawl"

// 참조
// https://github.com/kdama/gopl/tree/master/ch08/ex07

const out = "./out"

func main() {
	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- []string{"http://naver.com"} }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl.Crawl(link, out)
				}(link)
			}
		}
	}
}
