package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// binding from to json
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// queryString
	r.GET("/param", func(ctx *gin.Context) {
		// if not a parameter, it will return the default value
		username := ctx.DefaultQuery("username", "eastwin")
		address := ctx.Query("address")
		// 输出json
		ctx.JSON(200, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// form
	r.POST("/form", func(ctx *gin.Context) {
		username := ctx.DefaultPostForm("username", "eastwin")
		address := ctx.PostForm("address")
		// 输出json
		ctx.JSON(200, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// json
	r.POST("/json", func(ctx *gin.Context) {
		// 为了方便暂时忽略错误处理
		j, _ := ctx.GetRawData() // 获取原始的json数据
		// 定义map或结构体
		var m map[string]interface{}
		// 反序列化json数据到map
		_ = json.Unmarshal(j, &m)
		ctx.JSON(200, m)
	})

	// pathvariable
	r.GET("/user/:username", func(ctx *gin.Context) {
		username := ctx.Param("username")
		ctx.JSON(200, gin.H{
			"message":  "ok",
			"username": username,
		})
	})

	// binding to json
	r.POST(("/loginJson"), func(ctx *gin.Context) {
		var login Login

		if err := ctx.ShouldBind(&login); err == nil {
			fmt.Println("login info:", login)
			ctx.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

	})

	// 绑定form表单示例 (user=q1mi&password=123456)
	r.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
	r.GET("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	r.Run(":8080")
}
