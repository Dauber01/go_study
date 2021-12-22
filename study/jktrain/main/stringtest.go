package main

import (
	"fmt"
)

func main() {
	var ben = &Ben{id: 10, name: "Ben"}
	var Jerry = &Jerry{name: "Jerry"}
	var maker IceCreamMaker = ben

	var loop0, loop1 func()

	loop0 = func() {
		maker = ben
		go loop1()
	}

	loop1 = func() {
		maker = Jerry
		go loop0()
	}

	go loop0()

	for {
		maker.Hello()
	}

}

type IceCreamMaker interface {
	Hello()
}

//由于Ben和Jerry的内存布局不一致，所以在发生interface读写非原子性的时候，会报panic
type Ben struct {
	id   int
	name string
}

func (b *Ben) Hello() {
	fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
}

type Jerry struct {
	name string
}

func (j *Jerry) Hello() {
	fmt.Printf("Jerry says, \"Hello my name is %s\"\n", j.name)
}
