package answerques_test

import (
	"sync/atomic"
	"testing"
)

// 我们在关闭服务器的时候,让filter拒绝所有的新请求

type filter struct {
	// 处理请求的函数
	handler func(req interface{}) interface{}
	// 0代表不拒绝,不为0 代表拒绝
	reject int32
}

func (f *filter) handle(req interface{}) (interface{}, bool) {
	if atomic.LoadInt32(&f.reject) > 0 {
		return nil, false
	}
	return f.handler(req), true
}

func (f *filter) RejectNewRequest() {
	atomic.StoreInt32(&f.reject, 1)
}

func TestFilter(t *testing.T) {
	f := &filter{}
	f.handle(nil)
}
