package main

import (
	"github.com/gin-gonic/gin"
	"unit-cvt/internal/service"
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
		//exchange := api.Group("/exchange")
		{
			// 路径：GET /api/v1/exchange/latest?base=USD
			//exchange.GET("/latest", )
			// 路径：POST /api/v1/exchange/convert
			//exchange.POST("/convert", )
		}

		units := api.Group("/units")
		{
			// POST /api/v1/units/length
			units.POST("/length", service.ConvertLength)
		}
	}

	r.Run(":8080")
}