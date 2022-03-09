package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//https://godoc.org/gopkg.in/go-playground/validator.v8 有go用于验证的所有规则及表达式示例

type Persion struct {
	// binding中的即为条件,当需要同时满足的时候,用","来进行间隔;当是或的关系的时候,使用"|"
	Name string `form:"name" binding:"required"`
	//gt表示大于
	Age     int    `form:"age" binding:"required,gt=10"`
	Address string `form:"address" binding:"required"`
}

//curl -X GET "http://127.0.0.1:8080/test?name=wang&address=wuhan&age=9"
func main() {
	r := gin.Default()
	r.GET("/test", func(ctx *gin.Context) {
		var persion Persion
		//校验参数的先决条件在于先要对参数与结构体进行绑定
		err := ctx.ShouldBind(&persion)
		if err != nil {
			ctx.String(http.StatusBadRequest, "the param not valid, err: %V", err)
			//如果下面不加return的话,函数仍然会继续执行
			return
		}
		ctx.String(http.StatusOK, "persion is : %v", persion)
	})
	r.Run()
}
