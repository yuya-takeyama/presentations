package main

import (
	"bytes"
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
