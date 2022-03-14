package synctest

import (
	"fmt"
	"sync"
	"testing"
)

// sync.Map 并发安全map的使用方法
func TestSyncMap(t *testing.T) {
	m := sync.Map{}
	m.Store("cat", "tom")
	m.Store("mouse", "jerry")

	val, ok := m.Load("cat")
	if ok {
		// 断言,同时进行弱类型的转换
		fmt.Println(len(val.(string)))
	}
	// 可以使用强类型的转换
	str := "哈哈哈"
	data := []byte(str)
	fmt.Println(data)
}

// 测试sync的lock,不可重入
func TestSyncLock(t *testing.T) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	// 会发生死锁,只有一个goroutine,会导致程序崩溃,30s超时
	mutex.Lock()
	defer mutex.Unlock()
}

// 测试sync的lock,不可升级
func TestSyncRwLock(t *testing.T) {
	var mutex sync.RWMutex
	mutex.RLock()
	defer mutex.RUnlock()

	// 这句会导致死锁,30s超时
	mutex.Lock()
	defer mutex.Unlock()
}

// 测试只输出一次,变量要在全局的范围内声明
var once sync.Once

func TestOnce(t *testing.T) {
	for i := 0; i < 10; i++ {
		doOnce()
	}
}

func doOnce() {
	once.Do(func() {
		fmt.Println("只输出一次")
	})
}
