package main

import (
	"github.com/FoyonaCZY/QweeBlog/pkg/config"
	"github.com/FoyonaCZY/QweeBlog/routers"
	"strconv"
)

func main() {
	//读取配置文件
	config.Init()

	//初始化数据库

	//初始化路由配置
	server := routers.InitRouter()
	err := server.Run(":" + strconv.Itoa(config.Configs.Server.Port))
	if err != nil {
		return
	}
}
