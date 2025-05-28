package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Glob是通配符模式
	r.LoadHTMLGlob("**/*.html")
	r.GET("/posts/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "Posts Index",
		})
	})

	r.GET("/users/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "Users Index",
		})
	})

	r.Run(":8080")
}
