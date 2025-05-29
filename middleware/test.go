package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "eastwind") // 可以通过c.Set在请求上下文中设置值,后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 计算处理时间
		cost := time.Since(start)
		log.Println("cost:", cost)
	}
}
