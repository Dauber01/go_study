package project_tmp_test

import (
	"testing"
	"fmt"
)

type Option interface{
	Hello() error
	Bye() error
}

type Demo struct{

}

func (d *Demo)Hello() error {
	fmt.Println("hello")
	return nil
}

func (d *Demo)Bye() error {
	fmt.Println("Bye")
	return nil
}

type ShowDemo struct{
	Demo
	count int
}

func ArgsType(opthions ...Option){
	for _, op := range opthions{
		op.Hello()
	}
}

// 可以看到对于 ...Option 类型的参数,它的实现 struct 的包装 struct 也可以作为参数
func TestArgsType(t *testing.T){
	sd := &ShowDemo{}
	ArgsType(sd)
}
