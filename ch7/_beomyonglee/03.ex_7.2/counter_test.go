package counter

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	b := &bytes.Buffer{}
	c, n := CountingWriter(b)
	data := []byte("hi there")
	c.Write(data)
	if *n != int64(len(data)) {
		t.Logf("%d != %d", n, len(data))
		t.Fail()
	}
}
