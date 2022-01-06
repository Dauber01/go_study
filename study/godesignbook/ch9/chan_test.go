package ch9

//同一个目录中允许有 _test 结尾的两个package信息

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	tc := make(chan int, 10)
	for i := 0; i < 10; i++ {
		tc <- i
	}
	close(tc)

	for i := 0; i < 20; i++ {
		mm := <-tc
		//由测试可知,在 chan 被关闭,且消息消费完毕之后,继续接受仍不断接收到零值
		fmt.Println(mm)
	}
	//Urls()
	//tt := urls
	//fmt.Println(tt)
}
