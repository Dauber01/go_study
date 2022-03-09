package main

import (
	"testing"
	"time"
)

func TestPrint1(t *testing.T) {
	print1()
}

func TestGoPrint1(t *testing.T) {
	goPrint1()
	time.Sleep(1 * time.Millisecond)
}

func BenchmarkPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print1()
	}
}

func BenchmarkGoPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint1()
	}
}

//在go 1.5版本之前默认只会使用一个cpu来执行go的任务,但是在1.5之后,会用全部的cpu执行任务
//这时,需要使用 go test -run x -bench . -cpu 1 来只使用一个cpu
//go test -bench="." -cpu 1 最后用这个指令执行成功,所以每个gorountine的执行时间会是无限长么?
//goroutine对于每个任务的时限是有限制的么,机制是什么,基于什么原理或者算法来做的
func BenchmarkPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print2()
	}
}

func BenchmarkGoPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint2()
	}
}
