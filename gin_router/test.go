package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html") // 加载单个HTML文件
	// 一般路由
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 1,
		})
	})
	// 匹配所有方法
	r.Any("/test2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 2,
		})
	})
	// 404 method
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "index.html", nil)
	})

	// router group
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "User Index",
			})
		})
		userGroup.GET("/index2", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "User Index2",
			})
		})
	}
	r.Run(":8080") // listen and serve on
}
