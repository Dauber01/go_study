package answerques_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Limiter struct {
	// 当前处理请求的上限
	limit int32
	// 处理请求逻辑
	handler func(req interface{}) interface{}

	// 直接怼整个func进行锁定,不提倡这种写法,因为会将整个任务变为单线程,而且是使limit失去作用
	mutex sync.RWMutex

	// 用来计数当前已经开始的task数量
	cnt int32
}

// Reject bool 返回值表示究竟有没有执行,通过计数器来控制资源的应用,并判断task是否能够得到执行
func (l *Limiter) Reject(req interface{}) (interface{}, bool) {
	l.mutex.Lock()
	if l.cnt < l.limit {
		l.cnt++
		// 如果没到limit的数量,在对计数++保证可见性之后马上解锁
		l.mutex.Unlock()
		res := l.handler(req)
		// 在执行完后,为保证可见行和原子性,加锁,对cnt进行--
		l.mutex.Lock()
		l.cnt--
		return res, true
	}
	l.mutex.Unlock()
	return nil, false
}

// Reject bool (double check实现)
func (l *Limiter) DoubleCheckReject(req interface{}) (interface{}, bool) {
	l.mutex.RLock()
	// 快速失败,如果竞争紧张的时候,总可以进这个if的时候性能会好些
	if l.limit >= l.cnt {
		l.mutex.RUnlock()
		return nil, false
	}

	l.mutex.Lock()
	// double check 保证数据不会被覆盖
	if l.limit >= l.cnt {
		l.mutex.Unlock()
		return nil, false
	}
	l.cnt++
	l.mutex.Unlock()
	res := l.handler(req)
	l.mutex.Lock()
	l.cnt--
	defer l.mutex.Unlock()
	return res, true
}

// Reject bool (atomic double check实现)
func (l *Limiter) AtomicDoubleCheckReject(req interface{}) (interface{}, bool) {
	cnt := atomic.LoadInt32(&l.cnt)
	if cnt > l.limit {
		return nil, false
	}

	cnt = atomic.AddInt32(&l.cnt, 1)
	defer atomic.AddInt32(&l.cnt, -1)

	// 进行double check,用来对并发进入第二阶段的goroutine进行重新检查
	if cnt >= l.limit {
		return nil, false
	}
	return l.handler(req), true
}

// Reject bool (atomic no double check实现), 这个方法并没有对锁进行分段,
// 但因为整体的实现中都没有显示&全局的使用互斥锁,所以是可以接受的
func (l *Limiter) AtomicNoDoubleCheckReject(req interface{}) (interface{}, bool) {
	// 在进入方法的时候直接进行++,然后再进行判断
	cnt := atomic.AddInt32(&l.cnt, 1)
	defer atomic.AddInt32(&l.cnt, -1)
	// 发现超出上限之后直接进行返回
	if cnt > l.limit {
		return nil, false
	}
	return l.handler(req), true
}

func TestLimiter(t *testing.T) {

}
