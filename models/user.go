package models

import (
	"github.com/FoyonaCZY/QweeBlog/pkg/config"
	"github.com/FoyonaCZY/QweeBlog/util"
	"github.com/jinzhu/gorm"
)

const (
	UserStatusActive    = 1
	UserStatusNotActive = 2
	UserStatusBanned    = 3
)

// User 用户模型
type User struct {
	gorm.Model
	Nickname        string `gorm:"type:varchar(50);not null;" json:"nickname"`
	Email           string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password        string `gorm:"type:varchar(100);not null" json:"-"`
	Avatar          string `gorm:"type:varchar(100);not null" json:"avatar"`
	GroupID         uint   `gorm:"type:int;not null" json:"group_id"`
	ReceiveEmail    bool   `gorm:"type:boolean;not null" json:"receive_email"`
	Status          int    `gorm:"type:int;not null" json:"status"`
	ActivationToken string `gorm:"type:varchar(100);not null" json:"-"`

	// 关联模型
	Group Group `gorm:"save_associations:false:false"`
}

// GetUserByID 根据ID获取用户
func GetUserByID(id uint) (User, error) {
	var user User
	err := DB.Where("id = ?", id).Preload("Group").First(&user).Error
	return user, err
}

// GetUserByEmail 根据邮箱获取用户
func GetUserByEmail(email string) (User, error) {
	var user User
	err := DB.Where("email = ?", email).Preload("Group").First(&user).Error
	return user, err
}

// BeforeCreate 创建用户前的钩子
func (user *User) BeforeCreate() (err error) {
	return nil
}

// BeforeSave 保存用户前的钩子
func (user *User) BeforeSave() (err error) {
	return nil
}

// NewDefaultUser 新建默认用户
func NewDefaultUser() User {
	user := User{
		Email:           "",
		GroupID:         config.Configs.DefaultGroup.ID,
		Avatar:          config.Configs.DefaultAvatar.DefaultAvatar,
		ReceiveEmail:    true,
		Status:          UserStatusNotActive,
		ActivationToken: "nil",
	}
	if config.Configs.Smtp.Enable == 0 {
		user.Status = UserStatusActive
	}
	return user
}

func (user *User) SetPassword(password string) error {
	passwordHash, err := util.HashPassword(password)
	if err != nil {
		return err
	}
	user.Password = passwordHash
	return nil
}

func DeleteUserByID(ID uint) int64 {
	return DB.Delete(User{}, ID).RowsAffected
}

func GetAllUsers() ([]User, error) {
	var users []User
	err := DB.Preload("Group").Find(&users).Error
	return users, err
}
