package current_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func Server(Addr string, handler http.Handler, stop <-chan struct{}) error {

	s := http.Server{
		Addr:    Addr,
		Handler: handler,
	}

	go func() {
		// 当有信号进入的时候,阻塞结束,执行shutdown操作
		<-stop
		s.Shutdown(context.Background())
	}()

	// 当服务停止的时候,将信号量进行返回
	return s.ListenAndServe()
}

func TestMain(t *testing.T) {
	// 创建两个队列,分别用于统计停止的情况和传递停止指令
	done := make(chan error, 2)
	stop := make(chan struct{})

	go func() {
		done <- Server("8080", nil, stop)
	}()

	go func() {
		done <- Server("8081", nil, stop)
	}()

	var stopped bool

	// 从done通道中获得结束的信息
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Printf("the err is, %v", err)
		}
		// 第一次进入的时候,发送停止信号,停止两个通道
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
