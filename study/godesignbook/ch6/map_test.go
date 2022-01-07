package ch6

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[string]string)
	m["hello"] = "word"
	fmt.Println(m)
	add(m)
	//map为引用类型,所以在传递到其它函数之后,修改的属性为同一地址,所以会影响调用者变量
	fmt.Println(m)
	set(m)
	fmt.Println(m)
}

func add(m map[string]string) {
	m["ping"] = "pong"
}

func set(m map[string]string) {
	m["ha"] = "he"
	//因为go的参数值传递为值传递,所以变量 m 存储的为地址值,变量 m 被替换无法影响调用者
	m = nil
}
