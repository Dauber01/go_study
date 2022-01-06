package ch8

import (
	"fmt"
	"testing"
)

/* 当 chan 用作 event 的时候,目的只是单纯的用作同步,我们可以使用 struct{} 元素类型来强调
尽管也可以用 bool 或 int 类型通道来做同样的事,因为 done <- 1 比 done <- struct{} 要短 */
func Test(t *testing.T) {
	naturals := make(chan int)
	//对于这种没有缓存的通道,是同步的,可以保证并发中的先后顺序
	squares := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			naturals <- i
		}
		/* 也可以不关闭 chan ,仅在需要通知通道关闭的时候才需要关闭,go的垃圾回收会通过可达分析进行收集
		而不是通过通道是否关闭 */
		close(naturals)
	}()
	go func() {
		for num := range naturals {
			squares <- num * num
		}
		close(squares)
	}()
	for {
		//对于 chan 是否关闭,可以通过接收第二个参数为 false 进行间接的判断
		num, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(num)
	}

}
