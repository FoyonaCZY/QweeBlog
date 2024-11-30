package middlewares

import (
	"github.com/FoyonaCZY/QweeBlog/pkg/config"
	"github.com/gin-gonic/gin"
	"time"
)

func MaxAllowed() gin.HandlerFunc {
	sem := make(chan struct{}, config.Configs.Server.MaxConns)
	return func(c *gin.Context) {
		var called, fulled bool
		defer func() {
			if called == false && fulled == false {
				<-sem
			}

			if err := recover(); err != nil {

			}
		}()

		select {
		case sem <- struct{}{}:
			c.Next()
			called = true
			<-sem
		case <-time.After(time.Duration(config.Configs.Server.TimeOut) * time.Millisecond):
			fulled = true
			c.AbortWithStatus(504)
		}
	}
}
