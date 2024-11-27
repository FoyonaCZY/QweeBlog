package config

import (
	"github.com/go-ini/ini"
)

type Config struct {
	// 服务端口
	ServerPort int `ini:"port"`
}

func InitConfig() *Config {
	config := &Config{}

	// 读取配置文件
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		panic("读取配置文件失败: " + err.Error())
	}

	config.ServerPort, err = cfg.Section("server").Key("port").Int()
	if err != nil {
		panic("读取配置文件失败: 端口号必须为整数")
	}

	return config
}
