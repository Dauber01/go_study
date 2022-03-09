package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		//c.String(200, "get")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})
	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})
	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})
	r.Run()
	//使用如下指令进行访问
	//curl -X GET "http://127.0.0.1:8080/get"
}
