package config

import (
	"github.com/FoyonaCZY/QweeBlog/util"
	"github.com/go-ini/ini"
)

// Config 配置文件
type Config struct {
	//站点配置
	Site struct {
		Domain string `ini:"domain"`
	}

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
		Charset         string `ini:"charset"`
		MaxOpenConns    int    `ini:"maxopenconns"`
		MaxIdleConns    int    `ini:"maxidleconns"`
		ConnMaxLifetime int    `ini:"connmaxlifetime"`
	}

	//Jwt配置
	Jwt struct {
		Secret        string `ini:"secret"`
		TokenLifeTime int    `ini:"tokenexptime"`
	}

	//SMTP配置
	Smtp struct {
		Host     string `ini:"host"`
		Port     int    `ini:"port"`
		User     string `ini:"user"`
		Password string `ini:"password"`
		Nickname string `ini:"nickname"`
		Enable   bool   `ini:"enable"`
	}

	//新用户默认组
	DefaultGroup struct {
		ID uint `ini:"id"`
	}

	//默认头像
	DefaultAvatar string `ini:"default"`
}

var Configs *Config

// Init 读取配置文件
func Init() {
	// 读取配置文件
	cfg, err := ini.Load("pkg/config/config.ini")
	if err != nil {
		util.Panic("打开配置文件失败: " + err.Error())
	}
	Configs = &Config{}

	// 读取Site配置
	err = cfg.Section("site").MapTo(&Configs.Site)
	if err != nil {
		util.Panic("解析配置文件失败: " + err.Error())
	}

	// 读取Server配置
	err = cfg.Section("server").MapTo(&Configs.Server)
	if err != nil {
		util.Panic("解析配置文件失败: " + err.Error())
	}

	// 读取MySQL配置
	err = cfg.Section("mysql").MapTo(&Configs.MySQL)
	if err != nil {
		util.Panic("解析配置文件失败: " + err.Error())
	}

	// 读取Jwt配置
	err = cfg.Section("jwt").MapTo(&Configs.Jwt)
	if err != nil {
		util.Panic("解析配置文件失败: " + err.Error())
	}

	// 读取SMTP配置
	err = cfg.Section("smtp").MapTo(&Configs.Smtp)
	if err != nil {
		util.Panic("解析配置文件失败: " + err.Error())
	}

	// 读取DefaultGroup配置
	err = cfg.Section("defaultgroup").MapTo(&Configs.DefaultGroup)
	if err != nil {
		util.Panic("解析配置文件失败: " + err.Error())
	}

	// 读取DefaultAvatar配置
	Configs.DefaultAvatar = cfg.Section("avatar").Key("default").String()
}

// UpdateConfig 更新配置文件
func UpdateConfig() {
	cfg, err := ini.Load("pkg/config/config.ini")
	if err != nil {
		util.Panic("修改配置文件失败: " + err.Error())
	}

	err = cfg.Section("site").ReflectFrom(&Configs.Site)
	if err != nil {
		util.Panic("修改配置文件失败: " + err.Error())
	}
	err = cfg.Section("server").ReflectFrom(&Configs.Server)
	if err != nil {
		util.Panic("修改配置文件失败: " + err.Error())
	}
	err = cfg.Section("mysql").ReflectFrom(&Configs.MySQL)
	if err != nil {
		util.Panic("修改配置文件失败: " + err.Error())
	}
	err = cfg.Section("jwt").ReflectFrom(&Configs.Jwt)
	if err != nil {
		util.Panic("修改配置文件失败: " + err.Error())
	}
	err = cfg.Section("smtp").ReflectFrom(&Configs.Smtp)
	if err != nil {
		util.Panic("修改配置文件失败: " + err.Error())
	}
	err = cfg.Section("defaultgroup").ReflectFrom(&Configs.DefaultGroup)
	if err != nil {
		util.Panic("修改配置文件失败: " + err.Error())
	}
	cfg.Section("avatar").Key("default").SetValue(Configs.DefaultAvatar)

	err = cfg.SaveTo("pkg/config/config.ini")
	if err != nil {
		util.Panic("修改配置文件失败: " + err.Error())
	}
}
