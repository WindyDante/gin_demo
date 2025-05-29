package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// http redirect
	r.GET("/test", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	// router redirect
	r.GET("/test2", func(ctx *gin.Context) {
		// 指定重定向的url
		ctx.Request.URL.Path = "/test3"
		r.HandleContext(ctx)
	})
	r.GET("/test3", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	r.Run(":8080")
}
