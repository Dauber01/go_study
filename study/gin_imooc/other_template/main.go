package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//本质上模版的设定还是基于go官网的go template,具体内容可以参考官网
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "who are you",
		})
	})
	r.Run()
}
