package current_test

import (
	"fmt"
	"testing"
)

type IceCreamMaker interface {
	Hello()
}

type Ben struct {
	name string
}

func (b *Ben) Hello() {
	fmt.Printf("Ben says, hello my name is %s\n", b.name)
}

type Jerry struct {
	name string
}

func (j *Jerry) Hello() {
	fmt.Printf("Jerry says, hello my name is %s\n", j.name)
}

// 有一定概率出现interface的type和data不一致的情况
func TestInterface(t *testing.T) {
	var b = &Ben{name: "ben"}
	var j = &Jerry{name: "jerry"}
	var maker IceCreamMaker = b

	var loop0, loop1 func()

	loop0 = func() {
		maker = b
		go loop1()
	}

	loop1 = func() {
		maker = j
		go loop0()
	}

	go loop0()

	for {
		maker.Hello()
	}
}
