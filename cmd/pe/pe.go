package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Determine the editor to use
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	// stdin -> temporary file
	tmp, err := ioutil.TempFile("", "pe")
	if err != nil {
		abort(err)
	}
	defer func() {
		tmp.Close()
		os.Remove(tmp.Name())
	}()
	_, err = io.Copy(tmp, os.Stdin)
	if err != nil {
		abort(err)
	}

	// Launch the editor
	tty, err := os.OpenFile("/dev/tty", syscall.O_RDWR, 0)
	if err != nil {
		abort(err)
	}
	cmd := exec.Command(editor, tmp.Name())
	cmd.Stdin = tty
	cmd.Stdout = tty
	err = cmd.Run()
	if err != nil {
		abort(err)
	}

	// temporary file -> stdout
	tmp2, err := os.Open(tmp.Name())
	if err != nil {
		abort(err)
	}
	defer tmp2.Close()
	_, err = io.Copy(os.Stdout, tmp2)
	if err != nil {
		abort(err)
	}
}

func abort(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
