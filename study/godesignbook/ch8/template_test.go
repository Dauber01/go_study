package ch8

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTemplate(t *testing.T) {
	start := time.Now()
	//使用阻塞chan花费 131760 nano second
	//sizes := make(chan int)
	//使用非阻塞chan花费 138915 nano second,可能由于计算的时间过快,过于集中导致
	sizes := make(chan int, 30)
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ans := quare(i)
			sizes <- ans
		}(i)
	}
	go func() {
		wg.Wait()
		close(sizes)
	}()
	res := make([]int, 0)
	for num := range sizes {
		res = append(res, num)
	}
	fmt.Println(len(res))
	fmt.Println(res)
	fmt.Printf("cost time %d", time.Since(start).Nanoseconds())
}

func quare(i int) int {
	return i * i
}
