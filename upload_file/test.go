package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 处理multipart forms提交文件时默认的内存是32 MiB
	// 通过MaxMultipartMemory可以设置允许上传的最大内存大小
	// r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(ctx *gin.Context) {
		// 单个文件
		if file, err := ctx.FormFile("f1"); err == nil {
			log.Println("upload file:", file.Filename)
			dst := fmt.Sprintf("E:/gin_demo/upload_file/%s", file.Filename)
			// 上传文件到指定的目录
			ctx.SaveUploadedFile(file, dst)
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": fmt.Sprintf("File %s uploaded successfully", file.Filename),
			})
		} else {
			ctx.String(http.StatusInternalServerError, "Upload file error: %s", err.Error())
			return
		}
	})
	r.POST("/uploadList", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["f1"]

		// 遍历文件列表
		for _, file := range files {
			log.Println("upload file:", file.Filename)
			dst := fmt.Sprintf("E:/gin_demo/upload_file/%s", file.Filename)
			// 上传文件到指定的目录
			ctx.SaveUploadedFile(file, dst)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": fmt.Sprintf("Uploaded %d files successfully", len(files)),
		})
	})
	r.Run(":8080")
}
