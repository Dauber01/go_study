package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//restful 路由 get函数
func helloworldGet(c *gin.Context) {
	c.String(http.StatusOK, "[gin]hello word in get!")
}

func helloworldPost(c *gin.Context) {
	c.String(http.StatusOK, "[gin]hello word in post!")
}

func fetch(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, fmt.Sprintf("id is : %s\n", id))
}

func action1(c *gin.Context) {
	c.String(http.StatusOK, "action1")
}

func action2(c *gin.Context) {
	c.String(http.StatusOK, "action2")
}

func main1() {
	router := gin.Default()

	//restful 路由
	router.GET("/restful", helloworldGet)
	router.POST("/restful", helloworldPost)

	//不支持正则表达式
	//提取path中的参数
	router.GET("/param/:id", fetch)

	//组路由
	milestoneV1 := router.Group("/v1")
	{
		milestoneV1.GET("/action1", action1)
		milestoneV1.GET("/action2", action2)
	}

	//服务启动
	router.Run("127.0.0.1:8080")
}
