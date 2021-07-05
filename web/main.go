package main

import (
	"github.com/gin-gonic/gin"
	c "web/controller"
)


func main() {
	router := gin.Default()
	router.Static("/home","view")

	// 路由匹配
	//router.GET("/", func(context *gin.Context) {
	//	context.Writer.WriteString("项目开始了")
	//
	//})

	router.GET("/api/v1.0/session",c.GetSession)

	router.GET("api/v1.0/imagecode/:uuid",c.GetImageCd)
	router.Run(":8081")
}


