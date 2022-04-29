package current_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var Wait sync.WaitGroup
var Counter int = 0

// 执行 go test -race race_detector_test.go 指令即可进行检测竞态条件,并标出并发读写各发生在哪行
func TestRace(t *testing.T) {
	for i := 0; i < 2; i++ {
		Wait.Add(1)
		go Routine(i)
	}
	Wait.Wait()
	fmt.Printf("the final counter is :%d", Counter)
}

func Routine(id int) {
	for i := 0; i < 2; i++ {
		// 下面对应的三行代码也可以直接替换为 Counter++
		// 同样,虽然在代码中之后一行,但是翻译成汇编之后,实际上是三行代码,也非原子操作
		value := Counter
		time.Sleep(1 * time.Nanosecond)
		value++
		Counter = value
	}
	Wait.Done()
}
