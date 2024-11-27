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
	Password     string `gorm:"type:varchar(100);not null" json:"password"`
	GroupID      uint   `gorm:"type:int;not null" json:"group_id"`
	ReceiveEmail bool   `gorm:"type:boolean;not null" json:"receive_email"`

	// 关联模型
	Group Group `gorm:"save_associations:false:false"`
}

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

// SetPassword 根据明文加密并设置用户密码
func (user *User) SetPassword(password string) error {
	passwordHash, err := util.HashPassword(password)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	return nil
}
