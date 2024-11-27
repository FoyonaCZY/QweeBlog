package config

import (
	"github.com/go-ini/ini"
)

// Config 配置文件
type Config struct {
	// 服务器配置
	Server struct {
		Port int `ini:"port"`
	}

	//MySQL配置
	MySQL struct {
		Host            string `ini:"host"`
		Port            int    `ini:"port"`
		Database        string `ini:"database"`
		Username        string `ini:"username"`
		Password        string `ini:"password"`
		MaxOpenConns    int    `ini:"maxopenconns"`
		MaxIdleConns    int    `ini:"maxidleconns"`
		ConnMaxLifetime int    `ini:"connmaxlifetime"`
	}
}

var Configs *Config

// Init 读取配置文件
func Init() {
	// 读取配置文件
	cfg, err := ini.Load("pkg/config/config.ini")
	if err != nil {
		panic("打开配置文件失败: " + err.Error())
	}
	Configs = &Config{}

	// 读取Server配置
	err = cfg.Section("server").MapTo(&Configs.Server)
	if err != nil {
		panic("解析配置文件失败: " + err.Error())
	}

	// 读取MySQL配置
	err = cfg.Section("mysql").MapTo(&Configs.MySQL)
	if err != nil {
		panic("解析配置文件失败: " + err.Error())
	}
}
