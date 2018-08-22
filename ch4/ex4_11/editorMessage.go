package main

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
)

// editorMessage calls EDITOR with a tmp file and returns the result.
func editorMessage() ([]byte, error) {

	// check for environment variable
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return nil, errors.New("warning: no EDITOR environment variable")
	}
	// create tmp file & close it
	f, err := ioutil.TempFile("", "github-message")
	if err != nil {
		return nil, err
	}
	fpath := f.Name()
	f.Close()
	// invoke editor
	cmd := exec.Command(editor, fpath)

	// https://stackoverflow.com/questions/12088138/trying-to-launch-an-external-editor-from-within-a-go-program
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	// wait for editor process to finish
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	// read contents of tmp file
	message, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	err = os.Remove(fpath)
	if err != nil {
		return nil, err
	}
	return message, nil
}
