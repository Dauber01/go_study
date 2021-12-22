package main

import "testing"

func TestReduce(t *testing.T) {
	//m := make(map[string]int)
	m := map[string]int{}
	m["a"] = 0
	addaelem(m)
	//经测试,两次输出都有 addaelem 方法中加入的元素,说明 map 结构在传递过程中走的其实是饮用传递
	t.Log(m)
}

func addaelem(m map[string]int) {
	//为map增加一个元素
	m["hello"] = 1
}
