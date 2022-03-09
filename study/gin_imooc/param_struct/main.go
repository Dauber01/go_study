package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Persion struct {
	//标识字段中使用 `form` 表示将post/get的参数直接转为struct类型
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)
	r.Run()
}

//curl -X GET "http://127.0.0.1:8080/testing?name=wang&address=wuhan&birthday=2022-02-01"
//curl -H "Content-Type:application/json" -X POST "http://127.0.0.1:8080/testing" -d '{"name":"wang", "address":"wuhan"}'
func testing(c *gin.Context) {
	var persion Persion
	//c.ShouldBind是根据请求content-type来做不同的binding操作
	err := c.ShouldBind(&persion)
	if err != nil {
		c.String(http.StatusOK, "persion bind err: %v", err)
		c.Abort()
	}
	c.String(http.StatusOK, "%v", persion)
}
