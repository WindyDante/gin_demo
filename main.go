package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个默认的路由引擎
	router := gin.Default()
	// get请求,路径ping
	// 响应json 状态码200
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := router.Run()
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
