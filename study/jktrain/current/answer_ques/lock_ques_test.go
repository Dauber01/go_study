package answerques_test

import (
	"sync"
	"testing"
)

type SafeMap struct {
	m     map[string]interface{}
	mutex sync.RWMutex
}

// 方法原来有值的时候,返回原来的值;如原来不存在值,则将newValue的值存入,并返回
func (s *SafeMap) LoadOrStore(key string, newValue interface{}) (val interface{}, loaded bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()

	if ok {
		return val, true
	}

	s.mutex.Lock()

	// double check部分;如没有,对于多线程并发的情况下,可有两条goroutin进入到该阶段,在第一个gorotine执行结束之后
	// 第二个gorotine也会拿到锁,并且执行完这部分代码,但会形成覆盖的问题
	val, ok = s.m[key]
	if ok {
		return val, true
	}

	defer s.mutex.Unlock()
	s.m[key] = newValue
	return newValue, false
}

func TestLockQues(t *testing.T) {

}
