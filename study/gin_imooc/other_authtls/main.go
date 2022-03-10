package main

import (
	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello test")
	})
	//1.首先生成一个密钥
	//2.将密钥推送给证书机构,获得一个私钥
	//3.在本地验证私钥
	//4.验证成功将私钥保存,之后用于加密
	//*因为目测我不需要这个功能,所以就没有进行测试
	autotls.Run(r, "wwww.itpp.tk")
}
