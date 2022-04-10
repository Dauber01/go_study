package error_test

import (
	"errors"
	"fmt"
	"testing"

	xerrors "github.com/pkg/errors"
)

var errMy = errors.New("my")

func test0() error {
	return xerrors.Wrapf(errMy, "test0 failed")
}

func test1() error {
	return test0()
}

func test2() error {
	return test1()
}

func TestMain(t *testing.T) {
	err := test2()
	fmt.Printf("main: \n%+v\n", err)
}
