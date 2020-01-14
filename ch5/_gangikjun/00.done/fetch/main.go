package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func main() {

}

// fetch 는 url을 다운로드하고
// 지역 파일의 이름과 길이를 반환한다.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
