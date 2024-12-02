package middlewares

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 CORS 相关的响应头
		c.Header("Access-Control-Allow-Origin", "*") // 允许所有域访问
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 处理预检请求（OPTIONS 请求）
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		// 执行后续处理
		c.Next()
	}
}
