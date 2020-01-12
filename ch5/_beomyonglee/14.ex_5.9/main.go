package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`\$\w+|\${\w+}`)

func expand(s string, f func(string) string) string {
	wrapper := func(s string) string {
		if strings.HasPrefix(s, "${") {
			s = s[2 : len(s)-1]
		} else {
			s = s[1:]
		}
		return f(s)
	}
	return pattern.ReplaceAllStringFunc(s, wrapper)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("ex:5.9: ")

	subs := make(map[string]string, 0)
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			pieces := strings.Split(arg, "=")
			if len(pieces) != 2 {
				fmt.Fprintln(os.Stderr, "usage: ex5.9 KEY=VAL ...")
				os.Exit(1)
			}
			k, v := pieces[0], pieces[1]
			subs[k] = v
		}
	}

	missing := make([]string, 0)
	used := make(map[string]bool, 0)

	f := func(s string) string {
		v, ok := subs[s]
		if !ok {
			missing = append(missing, s)
			return "$" + s
		}
		used[s] = true
		return v
	}

	b := &bytes.Buffer{}
	b.ReadFrom(os.Stdin)
	fmt.Print(expand(b.String(), f))

	unused := make([]string, 0)
	for k, _ := range subs {
		if !used[k] {
			unused = append(unused, k)
		}
	}

	if len(unused) > 0 {
		log.Printf("unused bindings: %s", strings.Join(unused, " "))
	}
	if len(missing) > 0 {
		log.Printf("missing bindings: %s", strings.Join(missing, " "))
	}
}

/*
go build main.go
cat text | ./main name=bylee
Hi there bylee.

How is $place? I hope you've been getting a lot of $activity in. Is $someone there? I'm abso$expletivelutely going to be there soon.ex:5.9: missing bindings: place activity someone expletive
*/
