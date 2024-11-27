package config

import (
	"github.com/go-ini/ini"
)

// Config 配置文件
type Config struct {
	// 服务端口
	ServerPort int `ini:"port"`

	//MySQL配置
	MySQL struct {
		Host     string `ini:"host"`
		Port     int    `ini:"port"`
		Database string `ini:"database"`
		Username string `ini:"username"`
		Password string `ini:"password"`
	}
}

var Configs *Config

// InitConfig 读取配置文件
func InitConfig() {
	// 读取配置文件
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		panic("读取配置文件失败: " + err.Error())
	}
	Configs = &Config{}

	// 读取Server配置
	Configs.ServerPort, err = cfg.Section("server").Key("port").Int()
	if err != nil {
		panic("读取配置文件失败: Server端口号必须为整数")
	}

	// 读取MySQL配置
	Configs.MySQL.Port, err = cfg.Section("mysql").Key("port").Int()
	if err != nil {
		panic("读取配置文件失败: MySQL端口号必须为整数")
	}
	Configs.MySQL.Host = cfg.Section("mysql").Key("host").String()
	Configs.MySQL.Database = cfg.Section("mysql").Key("database").String()
	Configs.MySQL.Username = cfg.Section("mysql").Key("username").String()
	Configs.MySQL.Password = cfg.Section("mysql").Key("password").String()
}
