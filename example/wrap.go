package main

import (
	"fmt"
	"github.com/pharosnet/errors"
	"io"
)

func main() {
	e := errors.Wrap(io.EOF)
	fmt.Printf("%+v\n", e)
}
