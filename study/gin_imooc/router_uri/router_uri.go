package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		//gin.H{} 表示返回map结构的数据
		c.JSON(200, gin.H{
			"name": c.Param("name"),
			"id":   c.Param("id"),
		})
	})
	r.Run()
}
