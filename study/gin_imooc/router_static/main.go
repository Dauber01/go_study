package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/asserts", "./asserts")
	r.StaticFS("/static", http.Dir("static"))
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.Run()
	//go build -o router_static && ./router_static
	//其中 -o 表示输出指定文件的名字
}
