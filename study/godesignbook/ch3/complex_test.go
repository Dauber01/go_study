package ch3

import (
	"fmt"
	"math/cmplx"
	"testing"
)

func Test(t *testing.T) {
	//可以通过 complex 函数创建复数
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	//也可以通过在数字后面 +i 和实部相加获得复数
	fmt.Println(1i * 1i)
	z := 1 + 2i
	fmt.Println(z)

	//可使用 cmolx 包中提供的函数对复数进行操作
	fmt.Println(cmplx.Sqrt(-1))
}
