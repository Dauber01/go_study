package main

import (
	"fmt"
	"testing"
)

func TestValueScope(t *testing.T) {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		//在该作用域的 x 变量被声明之前, x 为上一个作用域的变量值
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
	fmt.Println()
}
