package current_test

import (
	"fmt"
	"sync"
	"testing"
)

var so *singleton = &singleton{}

var once sync.Once
var s *singleton

func Factory() func() *singleton {
	s := &singleton{}
	fmt.Println("哈哈哈哈哈")
	return func() *singleton {
		return s
	}
}

type singleton struct {
}

func GetSingleInstance() *singleton {
	once.Do(func() {
		s = &singleton{}
	})
	return s
}

func TestSingleton(t *testing.T) {
	for i := 0; i < 100; i++ {
		fun := Factory()
		s := fun()
		t.Logf("%p", s)
	}

	/* for i := 0; i < 100; i++ {
		k := GetSingleInstance()
		//t.Log("%T", k)
		fmt.Printf("%p \n", k)
	} */

	/* for i := 0; i < 100; i++ {
		t.Log(&so)
	} */
}
