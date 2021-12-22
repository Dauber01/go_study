package main

import (
	"fmt"

	"github.com/astaxie/beego"
)

//Restful controller路由
type RESTfulController struct {
	beego.Controller
}

func (this *RESTfulController) Get() {
	this.Ctx.WriteString("hello word in get method!")
}

//beego通过不同的方法区分对应的 get 和 post区别
func (this *RESTfulController) Post() {
	this.Ctx.WriteString("hello word in post method!")
}

//正则路由
type RegExpController struct {
	beego.Controller
}

func (this *RegExpController) Get() {
	this.Ctx.WriteString(fmt.Sprintln("regexp mod"))
	id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString(fmt.Sprintf("id is : %s\n", id))

	//想获取所有的后缀,则使用 splat
	splat := this.Ctx.Input.Param(":splat")
	this.Ctx.WriteString(fmt.Sprintf("splat is : %s\n", splat))

	//文件的路径信息
	path := this.Ctx.Input.Param(":path")
	this.Ctx.WriteString(fmt.Sprintf("path is : %s\n", path))

	//文件名字格式信息
	ext := this.Ctx.Input.Param(":ext")
	this.Ctx.WriteString(fmt.Sprintf("ext is : %s\n", ext))
}

func main() {
	//Restful controller路由
	beego.Router("restful", &RESTfulController{})

	//正则路由
	beego.Router("regexp/?:id", &RegExpController{})
	//可通过定义参数对逻辑进行区分
	beego.Router("regexpnum/:id([1-9]+)", &RegExpController{})
	//所有非零的字符
	beego.Router("regexpstr/:id([\\w]+)", &RegExpController{})
	//前后后缀的匹配
	beego.Router("regexppre/abc:id([1-9]+)de", &RegExpController{})
	//获取所有的值
	beego.Router("regexpall/*", &RegExpController{})
	//获取文件的信息
	beego.Router("regexpfile/*.*", &RegExpController{})

	//start server
	beego.Run("127.0.0.1:8080")
}

//beego的session支持的后端存储: {file, mysql, redis}
