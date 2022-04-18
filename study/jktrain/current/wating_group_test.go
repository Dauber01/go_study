package current_test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

// 放在全局变量的位置去处理各个 goroutine 的计数,对整体的计数进行阻塞
type Tracker struct {
	wg sync.WaitGroup
}

func (t *Tracker) Shutdown(ctx context.Context) error {

	ch := make(chan struct{})

	go func() {
		// 开启异步线程,在异步线程执行完之后发送关闭chan的信号
		t.wg.Wait()
		close(ch)
	}()

	select {
	case <-ch:
		// 收到关闭完成的信号,直接退出
		return nil
	case <-ctx.Done():
		// 收到时间到了的信号,返回异常 time out
		return errors.New("time out")
	}
}

const timeout = 5 * time.Second

func TestWaitGroup(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var tr Tracker
	err := tr.Shutdown(ctx)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

}
