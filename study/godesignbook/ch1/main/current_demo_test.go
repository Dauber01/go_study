package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestGorutine(t *testing.T) {
	urlArray := []string{"www.baidu.com", "www.souhu.com"}
	start := time.Now()
	//因为没有长度,所以应该是阻塞的chan
	ch := make(chan string)
	//为每个请求创建一个 go routine
	for _, url := range urlArray {
		go fetch(url, ch)
	}
	for i := 0; i < len(urlArray); i++ {
		fmt.Println(<-ch)
	}
	//close(ch)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	prex := "http://"
	url = prex + url
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	//code := resp.StatusCode
	//fmt.Printf("the response is %d\n", code)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
