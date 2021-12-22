package current_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

//本demo用lock实现了copy on write，同时在会写的时候，用了atomic确保回写的操作是原子的
func TestAtomic(t *testing.T) {
	type Map map[string]string
	var m atomic.Value
	m.Store(make(Map))
	var mu sync.Mutex
	read := func(key string) (val string) {
		m1 := m.Load().(Map)
		return m1[key]
	}
	insert := func(key, value string) {
		mu.Lock()
		defer mu.Unlock()
		m1 := m.Load().(Map)
		m2 := make(Map)
		for k, v := range m1 {
			m2[k] = v
		}
		m2[key] = value
		m.Store(m2)
	}
	_, _ = read, insert
}
