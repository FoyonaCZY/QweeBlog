package middlewares

import (
	"github.com/FoyonaCZY/QweeBlog/pkg/config"
	"github.com/gin-gonic/gin"
	"time"
)

func MaxAllowed() gin.HandlerFunc {
	// 初始化 sem
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
			// 未达到限速并发上限，获取 token
			c.Next() // 转入其他中间件逻辑

			// 关键逻辑
			called = true // 如果其他中间件提前捕获 panic，下面代码还是会被执行
			<-sem
		case <-time.After(time.Duration(config.Configs.Server.TimeOut) * time.Millisecond):
			// 达到并发上限，且等待 timeoutMs 毫秒仍然未获取票据（说明其他请求未归还 token）
			fulled = true
			c.AbortWithStatus(504)
		}
	}
}
