package current_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

// 可以看到打出的数有一些五位不是连续的
func TestSyncAtomic(t *testing.T) {

	var atom atomic.Value

	cfg := &Config{}
	// 在经过store的存储之后,我们发现可以进行保持连续性,也就是说通过atomic进行替换之后可以保证其原子性
	// 同时也发现,在进行操作的时候,如果atomic.store()中存储的是cfg的时候,是无效的,怀疑是因为对应的是指针
	// 而指针在同一可见行的情况下,下面的数组内容还在进行更改,所以对应要操作变量
	atom.Store(cfg.a)
	go func() {
		i := 0
		for {
			i++
			// 在写入的过程中,其它goroutine对其进行访问,这样存在写入的中间态的情况
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4}
			atom.Store(cfg.a)
		}
	}()
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				fmt.Printf("%v\n", atom.Load())
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
