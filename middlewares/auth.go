package middlewares

import (
	"github.com/FoyonaCZY/QweeBlog/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
