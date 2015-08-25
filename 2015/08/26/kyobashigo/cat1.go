package main

import (
	"io"
	"os"
)

func main() {
	file, _ := os.Open(os.Args[1]) // io.Reader
	io.Copy(os.Stdout, file)
}
