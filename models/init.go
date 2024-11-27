package models

import (
	"fmt"
	"github.com/FoyonaCZY/QweeBlog/pkg/config"
	"github.com/FoyonaCZY/QweeBlog/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

func Init() {
	util.Info("初始化数据库连接...")
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Configs.MySQL.Username,
		config.Configs.MySQL.Password,
		config.Configs.MySQL.Host,
		config.Configs.MySQL.Port,
		config.Configs.MySQL.Database))
	if err != nil {
		util.Panic("连接数据库失败: " + err.Error())
	}

	// 设置数据库连接池
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(config.Configs.MySQL.ConnMaxLifetime))
	db.DB().SetMaxIdleConns(config.Configs.MySQL.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.Configs.MySQL.MaxOpenConns)

	DB = db

	DB.AutoMigrate(&User{}, &Post{}, &Comment{}, &Group{})

	util.Info("数据库连接成功")

	// 初始化数据
	//创建初始用户组
	addDefaultGroup()

	//创建初始管理员
	addDefaultUser()
}

func addDefaultGroup() {
	_, err := GetGroupByID(1)
	if gorm.IsRecordNotFoundError(err) {

		// 创建初始管理员用户组
		group := Group{}
		group.Name = DefaultGroupNameAdmin
		group.Type = GroupTypeAdmin
		err := DB.Create(&group).Error
		if err != nil {
			util.Panic("创建初始用户组失败")
		}

		// 创建初始用户用户组
		group = Group{}
		group.Name = DefaultGroupName
		group.Type = GroupTypeUser
		err = DB.Create(&group).Error
		if err != nil {
			util.Panic("创建初始用户组失败")
		}

		util.Info("创建初始用户组成功")
	}
}

// 添加初始管理员
func addDefaultUser() {
	_, err := GetUserByID(1)
	if gorm.IsRecordNotFoundError(err) {
		user := User{}
		user.Nickname = "admin"
		user.Email = "admin@iamczy.com"
		password := util.GenerateRandomString(16)
		err := user.SetPassword(password)
		if err != nil {
			util.Panic("设置初始管理员密码失败")
		}
		user.GroupID = 1
		user.ReceiveEmail = true
		err = DB.Create(&user).Error
		if err != nil {
			util.Panic("创建初始管理员失败")
		}

		util.Info(fmt.Sprintf("创建初始管理员用户成功\n初始邮箱: %s\n初始密码: %s", user.Email, password))
	}
}
