package models

import (
	"github.com/FoyonaCZY/QweeBlog/util"
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Nickname     string `gorm:"type:varchar(50);not null;" json:"nickname"`
	Email        string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password     string `gorm:"type:varchar(100);not null" json:"-"`
	Avatar       string `gorm:"type:varchar(100);not null" json:"avatar"`
	GroupID      uint   `gorm:"type:int;not null" json:"group_id"`
	ReceiveEmail bool   `gorm:"type:boolean;not null" json:"receive_email"`

	// 关联模型
	Group Group `gorm:"save_associations:false:false"`
}

var (
	DefaultAvatar = "https://z1.ax1x.com/2023/08"
)

// GetUserByID 根据ID获取用户
func GetUserByID(id uint) (User, error) {
	var user User
	err := DB.Where("id = ?", id).First(&user).Error
	return user, err
}

// GetUserByEmail 根据邮箱获取用户
func GetUserByEmail(email string) (User, error) {
	var user User
	err := DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// BeforeCreate 创建用户前的钩子, 对密码进行哈希
func (user *User) BeforeCreate(*gorm.DB) (err error) {
	passwordHash, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = passwordHash
	return nil
}

// BeforeSave 保存用户前的钩子, 对密码进行哈希
func (user *User) BeforeSave(*gorm.DB) (err error) {
	passwordHash, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = passwordHash
	return nil
}

// NewDefaultUser 新建默认用户
func NewDefaultUser() User {
	return User{
		GroupID:      NewUserDefaultGroup,
		Avatar:       DefaultAvatar,
		ReceiveEmail: true,
	}
}
