package ch7

import (
	"fmt"
	"testing"
)

type Number struct {
	val int
}

type Inter interface {
	String() string
}

func (n Number) String() string {
	return fmt.Sprintf("%d\n", n.val)
}

func TestAssert(t *testing.T) {
	defer fmt.Println("---------end---------")
	var num Inter
	//可以看到强转在失败的时候,会直接报panic
	num = num.(Number)
	fmt.Println("哈哈哈哈哈")
	//转换失败的情况
	/* if num, ok := num.(Number); ok {
		fmt.Println(num.String())
	} else {
		fmt.Println(fmt.Errorf("转化失败, %T", num))
	} */
	//转换成功的demo
	/* num = Number{3}
	if num, ok := num.(Inter); ok {
		fmt.Println(num.String())
	} else {
		fmt.Errorf("转化失败, %T", num)
	} */
}
