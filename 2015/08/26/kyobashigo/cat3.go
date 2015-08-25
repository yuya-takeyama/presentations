package main

import (
	"io"
	"os"

	"github.com/yuya-takeyama/argf"
)

func main() {
	reader, _ := argf.Argf() // io.Reader
	io.Copy(os.Stdout, reader)
}
