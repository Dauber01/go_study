package ch8

import (
	"fmt"
	"testing"
	"time"
)

func TestSelete(t *testing.T) {
	//abort := make(chan int, 3)
	abort := make(chan int)
	go chanProdu(abort)
	//time.Sleep(time.Second * 10)
	select {
	//有 default 的时候为非阻塞
	case <-abort:
		fmt.Println("the demo run")
	/* 可以看到在有default的情况下,无论 chan 是有缓存的还是没有缓存的,当执行到
	case 时,如果 chan 中没有数据,则直接执行default */
	default:
		fmt.Println("the case go")
	}
}

func chanProdu(out chan<- int) {
	defer close(out)
	time.Sleep(time.Second * 5)
	out <- 3
}
