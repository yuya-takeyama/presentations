package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	Double(os.Stdin, os.Stdout)
}

func Double(stdin io.Reader, stdout io.Writer) {
	buf, _ := ioutil.ReadAll(stdin)

	stdout.Write(buf)
	stdout.Write(buf)
}
