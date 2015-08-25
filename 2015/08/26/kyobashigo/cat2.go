package main

import (
	"io"
	"os"
)

func main() {
	file1, _ := os.Open(os.Args[1])        // io.Reader
	file2, _ := os.Open(os.Args[2])        // io.Reader
	reader := io.MultiReader(file1, file2) // io.Reader
	io.Copy(os.Stdout, reader)
}
