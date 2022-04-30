package current_test

import (
	"sync"
	"testing"
	"time"
)

func TestBufferChan(t *testing.T) {

	// 通过输出可以看出有缓冲通道在缓冲未满的情况下不会阻塞通道
	c := make(chan string, 2)
	// 无缓冲的通道会进行阻塞
	//c := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		c <- `foo`
		println(`second pro`)
		c <- `bar`
		println(`product end`)
	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Second * 1)
		println(`cosumer begin`)
		println(`message: ` + <-c)
		println(`message: ` + <-c)
	}()

	wg.Wait()
}
