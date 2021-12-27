package ch5

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"golang.org/x/net/html"
)

//1.一般在关闭资源的时候使用 defer, 其关闭的往往在函数中成对出现
//2.可以有多个 defer, 按顺序执行
//3.对于会报错,而且需要记录关闭报错的时候,要谨慎使用 defer
func TestDefer(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		t.Log(err)
	}
	defer resp.Body.Close()
	ans, err := html.Parse(resp.Body)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(ans)
}

func TestTime(t *testing.T) {
	defer begin("hahah")()
	time.Sleep(5 * time.Second)
}

func begin(msg string) func() {
	//在开始的时候初始化函数,获得返回值
	start := time.Now()
	log.Printf("enter %s", msg)
	//在 defer 执行的时候,函数获取其外部变量,计算函数耗时
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
