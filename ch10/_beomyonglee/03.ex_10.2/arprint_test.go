package arprint_test

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/torbiak/gopl/ex10.2/tar"
	_ "github.com/torbiak/gopl/ex10.2/zip"
)

func TestOpen(t *testing.T) {
	for _, archive := range []string{"rah.zip", "rah.tar"} {
		b := &bytes.Buffer{}
		f, err := os.Open(filepath.Join("testdata", archive))
		if err != nil {
			t.Error(archive, err)
		}
		r, err := Open(f)
		if err != nil {
			t.Error(archive, err)
		}
		_, err = io.Copy(b, r)
		if err != nil {
			t.Error(archive, err)
		}
		want := `rah/b:
contentsB
rah/a:
contentsA
`
		got := b.String()
		if got != want {
			t.Errorf("os: got %q, want %q", archive, got, want)
		}
	}
}
