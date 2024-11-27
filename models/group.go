package models

import (
	"github.com/jinzhu/gorm"
)

const (
	GroupTypeAdmin        = 1
	GroupTypeUser         = 2
	DefaultGroupNameAdmin = "管理员"
	DefaultGroupName      = "普通用户"
)

var (
	// NewUserDefaultGroup 新用户默认组
	NewUserDefaultGroup int
)

// Group 用户组模型
type Group struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);not null;unique" json:"name"`
	Type uint   `gorm:"type:int;not null" json:"type"`
}

// GetGroupByID 根据ID获取用户组
func GetGroupByID(id uint) (Group, error) {
	var group Group
	err := DB.Where("id = ?", id).First(&group).Error
	return group, err
}
