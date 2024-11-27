package main

import (
	"github.com/FoyonaCZY/QweeBlog/config"
	"github.com/FoyonaCZY/QweeBlog/routers"
	"strconv"
)

func main() {
	//读取配置文件
	config.InitConfig()

	//初始化路由配置
	server := routers.InitRouter()
	err := server.Run(":" + strconv.Itoa(config.Configs.ServerPort))
	if err != nil {
		return
	}
}
