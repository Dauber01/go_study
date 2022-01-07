package ch6

import (
	"fmt"
	"testing"
	"time"
)

type Cia struct {
	name string
	age  int
}

func (c *Cia) Lanch() {
	fmt.Println("数据执行")
}

func TestFunc(t *testing.T) {
	c := Cia{"李雷", 18}
	//可以直接使用 c.Lanch 来标记函数
	time.AfterFunc(5*time.Second, c.Lanch)
	//睡眠8s,确定延迟执行的函数可以执行完毕
	time.Sleep(8 * time.Second)
}
