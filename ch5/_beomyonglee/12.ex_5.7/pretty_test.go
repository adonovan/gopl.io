package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettyOutputCanBeParsed(t *testing.T) {
	input := `
<html>
<body>
	<p class="something" id="short"><span class="special">hi</span></p><br/>
</body>
</html>
`
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Error(err)
	}
	pp := NewPrettyPrinter()
	b := &bytes.Buffer{}
	err = pp.Pretty(b, doc)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	_, err = html.Parse(bytes.NewReader(b.Bytes()))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
