package main

import (
	//用来操作mysql, redis等客户端
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gin_scaffold/router"

	//用来操作运行时环境
	"github.com/e421083458/golang_common/lib"
)

//https://gorm.io/zh_CN/docs 为gorm的文档,关于数据库映射
//https://godoc.org/github.com/gomodule/redigo/redis 为redigo的文档
func main() {
	lib.InitModule("./conf/dev", []string{
		"base", "mysql", "redis",
	})
	defer lib.Destroy()
	public.InitMysql()
	public.InitValidate()
	router.HttpServerRun()
}
