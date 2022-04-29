package current_test

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	//c := gen(2, 3)
	//out := sq(c)
	//fmt.Println(<-out)
	//fmt.Println(<-out)

	// 由于函数的参数和返回值类型相同,所以也可以写成如下形式
	for n := range sq(gen(2, 3)) {
		fmt.Println(n)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
			time.Sleep(time.Duration(time.Microseconds()) * 500)
			println("this")
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			time.Sleep(time.Second * 1)
			out <- n * n
		}
		close(out)
	}()
	return out
}
