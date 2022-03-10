package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

//也可以集成语言验证器,使校验信息和提示信息均为指定语言
type Booking struct {
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	//gtfield=CheckIn 表示要比字段的时间大
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

//由于我使用的包的版本问题,api已经发生变化
//https://godoc.org/gopkg.in/go-playground/validator.v10
func customFunc(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok {
		today := time.Now()
		if date.Unix() > today.Unix() {
			return true
		}
	}
	return false
}

//curl -X GET "http://127.0.0.1:8080/booking?check_in=2019-07-21&check_out=2022-09-09"
func main() {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", customFunc)
	}

	r.GET("/booking", func(ctx *gin.Context) {
		var b Booking
		err := ctx.ShouldBind(&b)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "valid not pass, e:" + err.Error() + ""})
			return
		}
		bookBytes, err := json.MarshalIndent(&b, "", "	")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "the struct not format"})
			return
		}
		ctx.JSON(http.StatusOK, string(bookBytes))
	})
	r.Run()
}
