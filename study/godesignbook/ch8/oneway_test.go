package ch8

import (
	"fmt"
	"testing"
)

func TestOneWay(t *testing.T) {
	defer fmt.Println("--------end---------")
	conters := make(chan int)
	squarers := make(chan int)
	//可以看到在使用中双向 chan 可以不受限制的被转换为单向
	go conter(conters)
	go squarer(conters, squarers)
	printNum(squarers)
}

func conter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	//但是反向装换是不允许的
	//out.(chan int)
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	//因为关闭的行为只能在写端操作,所以只有 chan<- 类型才可以关闭
	close(out)
}

func printNum(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
