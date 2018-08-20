package main

import (
	"fmt"
	"github.com/pharosnet/errors"
	"io"
)

func main() {
	e1 := io.EOF
	e2 := errors.With(e1, "error2")
	e3 := errors.WithF(e2, "%s", "error3")

	if errors.Contains(e3, e2) {
		// TODO ..
	}

	if errors.Contains(e3, e1) {
		// TODO ...
	}

	fmt.Println(errors.Contains(e3, e2))
	fmt.Println(errors.Contains(e3, e1))

}
