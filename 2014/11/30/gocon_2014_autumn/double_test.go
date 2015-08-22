package main

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestDouble(t *testing.T) {
	stdin := bytes.NewBufferString("hoge\n")
	stdout := new(bytes.Buffer)

	Double(stdin, stdout)

	if stdout.String() != "hoge\nhoge\n" {
		t.Fatal("Not matched")
	}
}

func TestCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "double.go")
	cmd.Stdin = bytes.NewBufferString("hoge\n")
	stdout := new(bytes.Buffer)
	cmd.Stdout = stdout

	_ = cmd.Run()

	if stdout.String() != "hoge\nhoge\n" {
		t.Fatal("Not matched")
	}
}
