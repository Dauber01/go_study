package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//对于所有以/user为前缀的都会被路由到这个函数中进行处理
	r.GET("/user/*action", func(ctx *gin.Context) {
		ctx.String(200, "success")
	})
	r.Run()
}
