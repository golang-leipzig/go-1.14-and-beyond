package main

import (
	"fmt"
	"io"
)

type ReadWriteCloser interface {
	io.ReadCloser
	io.WriteCloser
}

func main() {
	fmt.Println("works")
}
