package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func main() {
	r := gin.Default() // 返回默认路由引擎

	r.GET("/hello", sayHello)

	r.GET("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"method": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})

	// 启动服务
	r.Run(":8080") // 监听并在 8080 端口启动服务
}
