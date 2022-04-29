package current_test

import (
	"sync"
	"testing"
)

var so *singleton = &singleton{}

var once sync.Once
var s *singleton

// 这个看起来是单例的,是因为所有没有属性的struct的实例指针都是相同的
func Factory() func() *singleton {
	s := &singleton{}
	return func() *singleton {
		return s
	}
}

type singleton struct {
}

type singletonTest struct {
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
	for i := 0; i < 10; i++ {
		sk := &singleton{}
		t.Logf("%p", sk)
	}

	for i := 0; i < 10; i++ {
		st := &singletonTest{}
		t.Logf("%p", st)
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
