package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var sh = strings.TrimSpace(`
#!/bin/bash

/usr/bin/chromium --profile-directory=Default --app-id=hameamfeliimgoaebhdpioinjchdpkni
sleep 1
while (chromix-too ls | grep game.mahjongsoul.com 2>&1 >/dev/null)
  do sleep 5
done
`)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	tmp, err := ioutil.TempFile("", "discord-fake-process")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())

	tmp.WriteString(sh)
	tmp.Chmod(0755)
	tmp.Close()
	cmd := exec.Command(tmp.Name())
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Fprintln(os.Stderr, "ExitError:", exitErr)
			os.Exit(exitErr.ExitCode())
		}
		return err
	}

	return nil
}
