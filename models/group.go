package models

import (
	"github.com/jinzhu/gorm"
)

const (
	GroupTypeAdmin        = 1
	GroupTypeUser         = 2
	GroupTypeEditor       = 3
	DefaultGroupNameAdmin = "管理员"
	DefaultGroupNameUser  = "普通用户"
)

// GroupMetaData 用户组信息模型
type GroupMetaData struct {
	gorm.Model
	//默认用户组
	DefaultGroup uint `gorm:"type:int;not null" json:"default_group"`
}

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

// DeleteGroupByID 根据ID删除用户组
func DeleteGroupByID(id uint) int64 {
	return DB.Where("id = ?", id).Delete(Group{}).RowsAffected
}

// GetGroups 获取用户组列表
func GetGroups() ([]Group, error) {
	var groups []Group
	err := DB.Find(&groups).Error
	return groups, err
}

func GetNewUserDefaultGroup() uint {
	var groupMetaData GroupMetaData
	DB.First(&groupMetaData)
	return groupMetaData.DefaultGroup
}

func UpdateNewUserDefaultGroup(group uint) {
	var groupMetaData GroupMetaData
	DB.First(&groupMetaData)
	groupMetaData.DefaultGroup = group
	DB.Save(&groupMetaData)
}
