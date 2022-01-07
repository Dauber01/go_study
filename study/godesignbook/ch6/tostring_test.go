package ch6

import (
	"fmt"
	"testing"
)

type Address struct {
	province string
	city     string
}

func (a *Address) String() string {
	return fmt.Sprintf("pro : %s, city : %s", a.province, a.city)
}

func TestToString(t *testing.T) {
	a := Address{"北京", "朝阳"}
	fmt.Println(&a)
	fmt.Println(a.String())
	//fmt.Println() 会默认对变量的 .String() 方法进行输出,但是由于a的String()是绑定的指针,所以无法被加载到
	fmt.Println(a)
}
