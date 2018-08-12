/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// ex 4.11: Build a tool that lets users create, read, update, and close github
// issues from the command line, invoking their preferred text editor when sub-
// stantial text input is required.
package main

import "flag"

type options struct {
	search  bool
	create  bool
	read    bool
	update  bool
	close   bool
	repo    string
	issue   string
	message string // can also be search terms, space separated
}

var opts options

func init() {
	opts.create = *flag.Bool("create", false, "create issue")
}

func main() {
}
