package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//商品信息
type ProductInfo struct {
	Name  string
	Image string
	Price int64
}

func GetInfo() []ProductInfo {
	return []ProductInfo{
		{
			Name:  "lilei",
			Image: "/usr/images1.peng",
		},
		{
			Name:  "hanmeimei",
			Image: "/usr/hanmeimie.peng",
			Price: 16,
		},
	}
}

//返回 photos, title, prince 几项数据
func detailget(c *gin.Context) {
	id := c.Param("id")
	log.Printf("id is : %s\n", id)
	productInfoArray := GetInfo()
	c.JSON(http.StatusOK, gin.H{
		"hello":            "word",
		"productInfoArray": productInfoArray,
	})
}

func main2() {
	router := gin.Default()

	//设定html文件的请求目录,用来存放静态页面
	router.Static("/html", "./static")

	//路由,返回模拟的数据
	router.GET("/detail/:id", detailget)

	//服务启动
	router.Run("127.0.0.1:8080")
}
