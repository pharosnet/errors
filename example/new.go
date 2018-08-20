package main

import (
	"fmt"
	"github.com/pharosnet/errors"
)

func main() {
	e := errors.New("some error")
	fmt.Println(e)
	e = errors.ErrorF("%s", "2")
	fmt.Println(e)
	e = errors.NewByAssigned(32, 3, "new by assigned")
	fmt.Printf("%+v\n", e)
}
