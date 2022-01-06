package ch7

import (
	"fmt"
	"testing"
)

func TestSwich(t *testing.T) {

	fmt.Println(assertSwich(1))
	fmt.Println(assertSwich(false))
	fmt.Println(assertSwich("哈哈哈"))
	fmt.Println(assertSwich(3.43))
}

func assertSwich(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "nil"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		return "false"
	case string:
		return x
	default:
		panic(fmt.Sprintf("unsafe, %v %T\n", x, x))
	}
}
