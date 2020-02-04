package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"text/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token))
}

func (lex *lexer) eatWhitespace() int {
	i := 0
	for lex.token == ' ' || lex.token == '\n' {
		lex.next()
		i++
	}
	return i
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

type selector struct {
	tag   string
	attrs []attribute
}

func (s selector) String() string {
	b := &bytes.Buffer{}
	b.WriteString(s.tag)
	for _, attr := range s.attrs {
		switch attr.Value {
		case "":
			fmt.Fprintf(b, "[%s]", attr.Name)
		default:
			fmt.Fprintf(b, `[%s="%s"]`, attr.Name, attr.Value)
		}
	}
	return b.String()
}

type attribute struct {
	Name, Value string
}

func attrMatch(selAttrs []attribute, xmlAttrs []xml.Attr) bool {
SelectorAttribute:
	for _, sa := range selAttrs {
		for _, xa := range xmlAttrs {
			if sa.Name == xa.Name.Local && sa.Value == xa.Value || sa.Value == "" {
				continue SelectorAttribute
			}
		}
		return false
	}
	return true
}

func parseSelectors(input string) (_ []selector, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			panic(x)
		}
	}()

	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanStrings
	lex.scan.Whitespace = 0
	lex.next()

	selectors := make([]selector, 0)
	for lex.token != scanner.EOF {
		selectors = append(selectors, parseSelector(lex))
	}
	return selectors, nil
}

func parseSelector(lex *lexer) selector {
	var sel selector
	lex.eatWhitespace()
	if lex.token != '[' {
		if lex.token != scanner.Ident {
			panic(lexPanic(fmt.Sprintf("got %s, want ident", lex.describe())))
		}
		sel.tag = lex.text()
		lex.next()
	}
	for lex.token == '[' {
		sel.attrs = append(sel.attrs, parseAttr(lex))
	}
	return sel
}

func parseAttr(lex *lexer) attribute {
	var attr attribute
	lex.next()
	if lex.token != scanner.Ident {
		panic(lexPanic(fmt.Sprintf("got %s, want ident", lex.describe())))
	}
	attr.Name = lex.text()
	lex.next()
	if lex.token != '=' {
		if lex.token != ']' {
			panic(lexPanic(fmt.Sprintf("got %s, want ']'", lex.describe())))
		}
		lex.next()
		return attr
	}
	lex.next()
	switch lex.token {
	case scanner.Ident:
		attr.Value = lex.text()
	case scanner.String:
		attr.Value = strings.Trim(lex.text(), `"`)
	default:
		panic(lexPanic(fmt.Sprintf("got %s, want ident or string", lex.describe())))
	}
	lex.next()
	if lex.token != ']' {
		panic(lexPanic(fmt.Sprintf("got %s, want ']'", lex.describe())))
	}
	lex.next()
	return attr
}

func isSelected(stack []xml.StartElement, sels []selector) bool {
	if len(stack) < len(sels) {
		return false
	}
	start := len(stack) - len(sels)
	stack = stack[start:]
	for i := 0; i < len(sels); i++ {
		sel := sels[i]
		el := stack[i]
		if sel.tag != "" && sel.tag != el.Name.Local {
			return false
		}
		if !attrMatch(sel.attrs, el.Attr) {
			return false
		}
	}
	return true
}

func xmlselect(w io.Writer, r io.Reader, sels []selector) {
	dec := xml.NewDecoder(r)
	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if isSelected(stack, sels) {
				fmt.Fprintf(w, "%s\n", tok)
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	sels, err := parseSelectors(strings.Join(os.Args[2:], " "))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	xmlselect(os.Stdout, os.Stdin, sels)
}
