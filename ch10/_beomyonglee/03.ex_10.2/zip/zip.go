package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	arprint "github.com/torbiak/gopl/ex10.2"
)

type reader struct {
	zipReader *zip.Reader
	fileLeft  []*zip.File
	r         io.ReadCloser
	toWrite   string
}

func (r *reader) Read(b []byte) (int, error) {
	if r.r == nil && len(r.filesLeft) == 0 {
		return 0, io.EOF
	}
	if r.r == nil {
		f := r.filesLeft[0]
		r.filesLeft = r.filesLeft[1:]
		var err error
		r.r, err = r.Open()
		if err != nil {
			return 0, fmt.Errorf("read zip: %s", err)
		}
		if f.Mode()&os.ModeDir == 0 {
			r.toWrite = f.Name + ":\n"
		}
	}
	written := 0
	if len(r.toWrite) > 0 {
		n := copy(b, r.toWrite)
		b = b[n:]
		r.toWrite = r.toWrite[n:]
		written += n
	}
	n, err := r.r.Read(b)
	written += n
	if err != nil {
		r.r.Close()
		r.r = nil
		if err == io.EOF {
			return written, nil
		}
	}
	return written, nil
}

func NewReader(f *os.File) (io.Reader, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("new zip reader: %s", err)
	}
	r, err := zip.NewReader(f, stat.Size())
	if err != nil {
		return nil, fmt.Errorf("new zip reader: %s", err)
	}
	return &reader{r, r.File, nil, ""}, nil
}

func init() {
	arprint.RegisterFormat("zip", "PK", 0, NewReader)
}
