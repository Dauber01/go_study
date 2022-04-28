package current_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
两个goroutine获取锁差距巨大的主要原因在于:
1. 在开始的时候a首先获取锁,b发现锁住的时候进入队列等待锁
2. 在a执行完释放锁的时候,唤醒队列首部等待的b线程,b开始唤醒
3. 在b唤醒的过程中,a重新开始争抢锁
结论: 所以a很有可能在b唤醒的时候重新抢走锁,导致b唤醒后发现没有锁,只能重新进入队列开始等待
*/
func TestCompeteLock(t *testing.T) {
	done := make(chan bool, 1)
	var mu sync.Mutex

	countA := 0
	countB := 0

	// goroutine1 在测试中看到获取的次数为90次
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				time.Sleep(100 * time.Millisecond)
				countA++
				mu.Unlock()
			}
		}
	}()

	// goroutine2 在测试中看到获取的次数为10次
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			mu.Lock()
			countB++
			mu.Unlock()
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Printf("the case a' counter is : %d\n", countA)
	fmt.Printf("the case b' counter is : %d\n", countB)
}
