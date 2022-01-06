package ch9

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

var urls []string
var so sync.Once

func ArgsInit() {
	urls = []string{"http://www.baidu.com", "http://www.taobao.com", "http://www.jingdong.com", "http://www.taobao.com"}
}

func Urls() {
	so.Do(ArgsInit)
}

func TestPartLock(t *testing.T) {
	Urls()
	memo := New(httpGetBody)
	for _, url := range urls {
		go memo.Get(url)
	}
	time.Sleep(time.Second * 10)
	for k, v := range memo.cache {
		fmt.Printf("k:%s, v: %v, err: %v\n", k, v.res.value, v.res.err.Error())
	}
}

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type Func func(key string) (interface{}, error)

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}

func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	e := m.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		m.cache[key] = e
		m.mu.Unlock()
		e.res.value, e.res.err = m.f(key)
		close(e.ready)
	} else {
		m.mu.Unlock()
		//这个位置如果在通道关闭之后多次获取数据会有问题么
		<-e.ready
	}
	return e.res.value, e.res.err
}
