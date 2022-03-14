package atomictest

import (
	"sync/atomic"
	"testing"
)

var value int32 = 0

func TestAtomic(t *testing.T) {

	// 传入一个数字,并加10,输出
	atomic.AddInt32(&value, 10)
	num := atomic.LoadInt32(&value)
	println(num)

	// 当之前的值为10时换为20,输出结果 *CAS
	swapped := atomic.CompareAndSwapInt32(&value, 10, 20)
	println(swapped)

	// 将值替换为40,输出旧值
	old := atomic.SwapInt32(&value, 40)
	println(old)
	// 值变为最新交换的40
	println(value)
}
