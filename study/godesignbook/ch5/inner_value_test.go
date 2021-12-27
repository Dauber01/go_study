package ch5

import (
	"fmt"
	"testing"
)

func TestInnerValue(t *testing.T) {
	var funcs []func()
	for i := 0; i < 10; i++ {
		//使用内部变量tmp
		tmp := i
		/* 因为对于外面的迭代而言,i为在 for 域开始的时候初始的一个存储位
		i变量,迭代结束时,所有的变量会变为10 */
		funcs = append(funcs, func() {
			//fmt.Println(i * i)
			fmt.Println(tmp * tmp)
		})
	}
	for _, v := range funcs {
		v()
	}
}
