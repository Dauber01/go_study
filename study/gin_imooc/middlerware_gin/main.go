package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//实际发现不进行recover也是可以继续接受请求的,只是在日志存储上貌似有问题
//curl -X GET "http://127.0.0.1:8080/test?name=lilei"
func main() {

	//用来将默认日志和默认错误日志输入到文件
	//*可以发现,在输出到日志中之后,就不再在控制台打出了
	f, _ := os.Create("test.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	//我们在使用gin默认的初始化实例时,已经在方法内部初始化了log和recover方法
	//r := gin.Default()
	r := gin.New()
	//默认的logger会在请求进入的时候打出部分请求信息,已经响应状态和响应时长
	//在不使用recovery时,出现panic则会导致主线程无法继续响应请求
	r.Use(gin.Logger())
	//加入recovery之后,可以避免线程被结束
	//r.Use(gin.Logger(), gin.Recovery())
	r.GET("/test", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "hanmeimei")
		if name == "lilei" {
			panic("test panic")
		}
		ctx.String(http.StatusOK, "name:%s", name)
	})
	r.Run()
}
