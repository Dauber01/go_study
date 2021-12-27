package ch5

import (
	"fmt"
	"strings"
	"testing"
)

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

//代表每个字符的时候,只能用rune变量
func add(b rune) rune {
	return b + 1
}

func TestFunc(t *testing.T) {
	//函数可以用变量来表示
	f := square
	//同时传入参数,函数的结果也可以直接作为其他函数的参数
	fmt.Println(f(3))

	//同样将不同函数赋值给相同变量
	f = negative
	fmt.Println(f(3))
	// func(int) int 输出类型,可以看到f的类型为函数
	fmt.Printf("%T", f)

	//将函数product赋值给f的时候编译不能通过,所以看出函数具有类型
	//f = product

	//函数变量初始值为nil,所以可以和nil进行比较
	if f != nil {
		fmt.Println("可以看出,函数变量可以和 nil 进行比较")
	}

	//可以看出,不同的函数不能比较
	/* b := square
	if b == f {

	} */

	//strings.Map() 函数用来将字符串拆成字符,输入函数中,并将返回值组合到一起
	fmt.Println(strings.Map(add, "jaj"))
}
