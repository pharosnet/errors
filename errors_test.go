package errors_test

import (
	"fmt"
	"testing"

	"github.com/pharosnet/errors"
)

func TestErrorF(t *testing.T) {
	e1 := errors.New("error 1")
	e2 := errors.WithF(e1, "error %d", 2)
	fmt.Println(fmt.Sprintf("%+v", e2))

	e3 := errors.WithDepth(2, e2, "error 3")
	fmt.Println(fmt.Sprintf("%-v", e3))

}
