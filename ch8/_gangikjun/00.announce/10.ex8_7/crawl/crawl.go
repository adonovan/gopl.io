// Package crawl ..
package crawl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// 참조
// https://github.com/kdama/gopl/tree/master/ch08/ex07

var tokens = make(chan struct{}, 20)

// Crawl 대상페이지를 크롤링하여 로컬 디스크에 저장
func Crawl(path, out string) (nextPaths []string) {
	fmt.Println(path)

	tokens <- struct{}{} // acquire a token
	nextPaths, err := crawl(path, out)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}

	return nextPaths
}

func crawl(path, out string) (nextPaths []string, err error) {
	body, isHTML, err := fetch(path)
	if err != nil {
		return
	}

	var links []string
	if isHTML {
		links, body, err = extract(path, body)
		if err != nil {
			return
		}
	}

	err = Save(path, out, body)
	if err != nil {
		return
	}

	for _, extracted := range links {
		same, err := sameDomain(path, extracted)
		if err != nil {
			log.Print(err)
			continue
		}
		if same {
			nextPaths = append(nextPaths, extracted)
		}
	}
	return
}

// Save 는 주어진 바이트 배열 데이터를 로컬 디스크에 저장합니다.
func Save(path, out string, data []byte) error {
	parsed, err := url.Parse(path)
	if err != nil {
		return err
	}

	localPath := out + "/" + parsed.Host + parsed.Path
	if parsed.Path == "" || !strings.Contains(filepath.Base(parsed.Path), ".") {
		localPath += "/index.html"
	}
	err = os.MkdirAll(filepath.Dir(localPath), os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(localPath)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return err
}

// fetch path에서 필요한 것들을 가지고 온다
func fetch(path string) (body []byte, isHTML bool, err error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("getting %s: %s", path, resp.Status)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}

	isHTML = strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html") ||
		strings.HasPrefix(resp.Header.Get("Content-Type"), "text/xhtml")
	return body, isHTML, err
}

// extract 는 대상 url을 로컬 디스크에 저장, HTML 페이지 내의 링크 추출
// 가능하다면 해당 페이지 내의 링크를 상대경로로 변환
func extract(path string, data []byte) (links []string, converted []byte, err error) {
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, nil, fmt.Errorf("parsing %s as HTML: %v", path, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				pathURL, err := url.Parse(path)
				if err != nil {
					continue // ignore bad URLs
				}
				link, err := pathURL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
				n.Attr[i].Val = relativeURL(pathURL, link)
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	var b bytes.Buffer
	html.Render(&b, doc)
	return links, b.Bytes(), nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// base URL을 기점으로, link url을 상대 url로 변경
func relativeURL(base, link *url.URL) (result string) {
	if link.Scheme != base.Scheme {
		return link.String()
	}
	if link.Host != base.Host {
		return link.String()
	}
	depth := strings.Count(base.Path, "/") - 1
	if depth < 0 {
		depth = 0
	}
	result = "./" + strings.Repeat("../", depth) + link.Path
	if link.Path == "" || !strings.Contains(filepath.Base(link.Path), ".") {
		result += "/index.html"
	}
	if link.RawQuery != "" {
		result += "?" + link.RawQuery
	}

	re := regexp.MustCompile(`/+`)
	result = re.ReplaceAllString(result, "/")
	return
}

func sameDomain(x, y string) (bool, error) {
	px, err := url.Parse(x)
	if err != nil {
		return false, err
	}
	py, err := url.Parse(y)
	if err != nil {
		return false, err
	}
	return px.Host == py.Host, nil
}
