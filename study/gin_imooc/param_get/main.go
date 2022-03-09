package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//来获取param中的参数
	//curl -X GET "http://127.0.0.1:8080/param?first_name=wang"
	r.GET("/param", func(ctx *gin.Context) {
		firstName := ctx.Query("first_name")
		lastName := ctx.DefaultQuery("last_name", "lilei")
		ctx.String(http.StatusOK, "%s, %s", firstName, lastName)
	})

	//来获取body中的参数
	//curl -X POST "http://127.0.0.1:8080/body" -d '{"name":"wang"}'
	r.POST("/body", func(ctx *gin.Context) {
		//在Readall之后,ctx.PostForm()方法无法获取参数
		bodyData, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
		}
		//所以要在获取得对应的值之后在将其写回去
		//curl -X POST "http://127.0.0.1:8080/body" -d 'first_name=wang&last_name=gang'
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyData))
		firstName := ctx.PostForm("first_name")
		lastName := ctx.DefaultPostForm("last_name", "lilei")
		ctx.String(http.StatusOK, "firstName:%s, lastName:%s, content:%s",
			firstName, lastName, string(bodyData))
	})

	r.Run(":8080")
}
