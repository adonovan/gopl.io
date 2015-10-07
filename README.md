# The Go Programming Language

This repository provides the downloadable example programs
for the book, "The Go Programming Language"; see http://www.gopl.io.

(At present, only `gopl.io/ch1/helloworld` is available.  The other
programs will be published as soon as the book is printed.)

You can download, build, and run the programs with the following commands:

	$ export GOPATH=$HOME/gobook            # choose workspace directory
	$ go get gopl.io/ch1/helloworld         # fetch, build, install
	$ $GOPATH/bin/helloworld                # run
	Hello, 世界

Many of the programs contain comments of the form `//!+` and `//!-`.
These comments bracket the parts of the programs that are excerpted in the
book; you can safely ignore them.  In a few cases, programs
have been reformatted in an unnatural way so that they can be presented
in stages in the book.

