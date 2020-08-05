package errors_test

import (
	"fmt"
	"testing"

	"github.com/pharosnet/errors"
)

func TestErrorF(t *testing.T) {
	e1 := errors.New("error 1")
	e2 := errors.Withf(e1, "error %d", 2)
	fmt.Println(fmt.Sprintf("%+v", e2))

	e3 := errors.WithDepth(2, 3, e2, "error 3")
	fmt.Println(fmt.Sprintf("%-v", e3))

}

func wrap(msg string) error {
	return errors.NewWithDepth(1, 4, msg)
}

func Test_Wrap(t *testing.T) {

	e1 := wrap("1")
	e2 := errors.With(e1, "error 2")
	e3 := errors.With(e2, "error 3")

	fmt.Println(fmt.Sprintf("%+v", e3))

}
