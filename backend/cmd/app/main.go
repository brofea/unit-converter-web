package main

import (
	"unit-cvt/internal/service"

	"github.com/gin-gonic/gin"
)

// 简单的 CORS 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	// 添加 CORS 中间件
	r.Use(CORSMiddleware())

	// API 分组
	api := r.Group("/api/v1")
	{
		api.POST("/convert", service.Convert)
	}

	r.Run(":8080")
}
