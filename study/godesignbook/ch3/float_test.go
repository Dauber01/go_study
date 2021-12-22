package ch3

import (
	"fmt"
	"testing"
)

func TestFolatRange(t *testing.T) {
	var f float32 = 1 << 24
	/* float32 有效数字大约是6位, float64 有效数字大约15位,
	绝大多数情况应该使用 float64, float32 表示的正整数范围小容易出现误差 */
	fmt.Println(f == f+1)
}
