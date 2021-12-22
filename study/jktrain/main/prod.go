package main

import (
	"context"
	"fmt"
	"net/http"
)

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func serveApp(stop chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	return serve("0.0.0.0:8080", mux, stop)
}

func serveDebug(stop chan struct{}) error {
	return serve("127.0.0.1:8081", http.DefaultServeMux, stop)
}

func main2() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- serveDebug(stop)
	}()
	go func() {
		done <- serveApp(stop)
	}()

	var stopped bool
	//当有chan done执行完毕之后，就会执行关闭stop chan，这样两个server就都会被关闭
	//done设置长度是因为怕发生内存泄漏，而stop不设置长度因为直接调用close方法进行通讯，不会发生阻塞
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("error: %v", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
