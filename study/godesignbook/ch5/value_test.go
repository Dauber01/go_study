package ch5

import (
	"fmt"
	"testing"
)

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func TestValue(t *testing.T) {

	s := squares()
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
}
