package main

import "github.com/gin-gonic/gin"

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// json渲染
	users := User{
		Name:  "John Doe",
		Age:   30,
		Email: "12345"}
	r := gin.Default()
	r.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, users)
	})
	// gin.H是map[string]interface的缩写
	r.GET("/userT", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
			"code":    200,
		})
	})
	// xml渲染
	r.GET("/userXml", func(ctx *gin.Context) {
		ctx.XML(200, users)
	})
	// gin.H是map[string]interface的缩写
	r.GET("/userXmlT", func(ctx *gin.Context) {
		ctx.XML(200, gin.H{
			"message": "Hello, World!",
			"code":    200,
		})
	})
	// yaml渲染
	r.GET("/userYaml", func(ctx *gin.Context) {
		ctx.YAML(200, users)
	})

	r.Run(":8080")
}
