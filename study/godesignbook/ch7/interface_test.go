package ch7

import (
	"bytes"
	"io"
	"testing"
)

//利用 _ 变量和 nil 值进行断言,可以省掉2个变量
var _ io.Writer = (*bytes.Buffer)(nil)

func TestInterface(t *testing.T) {

}
