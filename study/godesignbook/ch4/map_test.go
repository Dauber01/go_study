package ch4

import (
	"fmt"
	"log"
	"testing"
)

func TestMap(t *testing.T) {
	ans := map[string]string{}
	ans["hello"] = "word"
	ca := &ans
	sout(ans)
	soutPoint(ca)
	log.Printf("%T", ca)
	log.Printf("%T", ans)
	log.Printf("%v", ans)

	array := make([]int, 2)
	array[0] = 1
	ap := &array
	log.Println(array)
	soutArray(array)
	//开始迷茫,到底传的是啥
	log.Println(array)
	log.Println(ap)
}

func soutArray(array []int) {
	//对非指针的操作也会生效
	array[1] = 2
}

func sout(ans map[string]string) {
	fmt.Println(ans)
	//对非指针的操作也会生效
	ans["go"] = "go"
}

func soutPoint(ca *map[string]string) {
	fmt.Println(ca)
	//指针也可以反获取
	(*ca)["point"] = "point"
}
