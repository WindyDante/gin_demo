package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 设置自定义模板函数
	router.SetFuncMap(template.FuncMap{
		// "safe" 函数将字符串转换为 HTML 安全的格式
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 加载模板文件
	router.LoadHTMLFiles("./index.tmpl")
	router.GET("/index", func(ctx *gin.Context) {
		// 使用自定义模板函数 "safe" 来处理字符串
		ctx.HTML(http.StatusOK, "index.tmpl", "<a href='https://1wind.cn'>eastwind blog</a>")
	})

	router.Run(":8080")
}
