package reader

import (
	"io"
)

type limitReader struct {
	r        io.Reader
	n, limit int
}

func (r *limitReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p[:r.limit])
	r.n += n
	if r.n >= r.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &limitReader{r: r, limit: limit}
}
