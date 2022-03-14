package error_test

import (
	"errors"
	"fmt"
	"testing"
)

type MyError struct {
}

func (m *MyError) Error() string {
	return "hello, this is my error"
}

// 测试 error 的一些用法
func TestErrors(t *testing.T) {
	var error MyError
	// panic(error.Error())
	se := errors.New("some error")
	// 初次进行对比的时候,发现对比的结果不是一种类型
	compare := errors.Is(&error, se)
	t.Log(compare)
	errors.As(&error, &se)
	t.Logf("%s", error.Error())
	t.Logf("%v", se)
	// 发现在经过类型转换之后,se的类型转变为MyError类型
	compare = errors.Is(&error, se)
	t.Log(compare)
}

// 测试在 panic 之后 recover的情况
func TestPanic(t *testing.T) {
	defer func() {
		// 可以看到在 panic发生之后,data内有数据,会打印出 panic的内容
		if data := recover(); data != nil {
			fmt.Printf("this is a panic: %v\n", data)
		}
		// panic之后,defer依然会执行,不会再执行pannic后的内容
		fmt.Println("恢复之后从这里开始执行")
	}()
	panic("哈哈哈哈")
	// 如提示 panic之后后续的内容不会继续执行
	//fmt.Println("这里将不会执行")
}

// 闭包特性
func TestFun(t *testing.T) {
	i := 13
	a := func() {
		fmt.Printf("i is %d \n", i)
	}
	// 在调用的时候可以看到输出
	a()
}

// 闭包引用作用域外变量
func TestDelay(t *testing.T) {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			// 在闭包内使用闭包外的参数,值在最终调用的时候才确定下来
			fmt.Printf("the number is %d\n", i)
		})
	}
	for _, fn := range fns {
		fn()
	}
}
