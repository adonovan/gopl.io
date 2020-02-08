// 메모테스트는 뭘 하는 걸까?

package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}


var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string) 
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		}{
			ch<- url
		}
		close(ch)
	}() 
	return ch
}


type M interface{
	Get(key string) (interface{}, error)
}


func Sequential(t *testing.T, m M) {
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", 
			url, time.Since(start), len(value.([]byte)))
	}
}


func Concurrent(t *testing.T, m M) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n", 
				url, time.Since(start), len(value.([]byte)))
		}(url)
		n.Wait()
	}
}