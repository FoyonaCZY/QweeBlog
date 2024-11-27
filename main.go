package main

import (
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/pkg/config"
	"github.com/FoyonaCZY/QweeBlog/routers"
	"github.com/FoyonaCZY/QweeBlog/util"
	"strconv"
)

func main() {
	//初始化日志
	util.InitLogger()

	//读取配置文件
	config.Init()

	//初始化数据库
	models.Init()

	//初始化路由配置
	server := routers.InitRouter()
	util.INFO("Server is running on port: " + strconv.Itoa(config.Configs.Server.Port))

	util.INFO("初始化完成")
	//启动服务
	err := server.Run(":" + strconv.Itoa(config.Configs.Server.Port))
	if err != nil {
		return
	}
}
