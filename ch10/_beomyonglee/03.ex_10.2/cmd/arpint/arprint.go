package main

import (
	"fmt"
	"io"
	"log"
	"os"

	arprint "github.com/torbiak/gopl/ex10.2"
	_ "github.com/torbiak/gopl/ex10.2/tar"
	_ "github.com/torbiak/gopl/ex10.2/zip"
)

func printArchive(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := arprint.Open(f)
	if err != nil {
		return fmt.Errorf("open archive reader: %s", err)
	}
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		return fmt.Errorf("printing: %s", err)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: arprint FILE ...")
	}
	exitCode := 0
	for _, filename := range os.Args[1:] {
		err := printArchive(filename)
		if err != nil {
			log.Print(err)
			exitCode = 2
		}
	}
	os.Exit(exitCode)
}
