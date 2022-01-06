package ch9

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

//使用 go test -race -v trace_test.go 指令进行执行的时候,发现会报data race的错误,并指明出现错误的行数
func TestTrace(t *testing.T) {
	fmt.Println("----------begin----------")
	num := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			num = i
			println(num)
		}(i)
	}
	wg.Wait()
	fmt.Println("----------end----------")
}
