package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ipwhitelist() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ipList := []string{
			//"127.0.0.1",
			"127.0.0.2",
		}
		clientIp := ctx.ClientIP()
		mark := false
		for _, ip := range ipList {
			if ip == clientIp {
				mark = true
				break
			}
		}
		if !mark {
			ctx.String(http.StatusOK, "%s is not in list", clientIp)
			ctx.Abort()
		}
	}
}

//curl -X GET "http://127.0.0.1:8080/test?name=lilei"
func main() {
	r := gin.Default()
	//可以通过 r.Use() 方法,向框架中插入对应的中间件执行功能
	r.Use(Ipwhitelist())
	r.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello test")
	})
	r.Run()
}
