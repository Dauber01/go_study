package ch5

import (
	"fmt"
	"testing"
)

func TestRecover(t *testing.T) {
	type layOut struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			fmt.Println("没有宕机")
		case layOut{}:
			fmt.Println("预期到的错误")
		default:
			panic(p)
		}
	}()
	for i := 0; i < 10; i++ {
		if i == 11 {
			panic(layOut{})
		}
	}
}
