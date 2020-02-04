package reader

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	s := "hi there"
	b := &bytes.Buffer{}
	r := LimitReader(strings.NewReader(s), 4)
	n, _ := b.ReadFrom(r)
	if n != 4 {
		t.Logf("n=%d", n)
		t.Fail()
	}
	if b.String() != "hi t" {
		t.Logf(`"%s" != "%s"`, b.String(), s)
	}
}
