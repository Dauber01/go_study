package ch7

import (
	"bytes"
	"io"
	"testing"
)

//利用 _ 变量和 nil 值进行断言,可以省掉2个变量
var _ io.Writer = (*bytes.Buffer)(nil)

func TestInterface(t *testing.T) {
	contrys := make(map[string]string)
	contrys["china"] = "中国"
	contrys["english"] = "英国"
	contrys["france"] = "法国"
	for k, v := range contrys {
		t.Logf("%s=%s ,", k, v)
	}
	t.Logf("变得聪明不是挺好的选择么")
}
