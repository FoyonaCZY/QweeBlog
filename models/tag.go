package models

import (
	"github.com/jinzhu/gorm"
)

// Tag 标签模型
type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);not null;unique" json:"name"`

	// 关联模型
	Posts []Post `gorm:"many2many:post_tags;save_associations:false:false"`
}

// GetTagByID 根据ID获取标签
func GetTagByID(id uint) (Tag, error) {
	var tag Tag
	err := DB.Where("id = ?", id).First(&tag).Error
	return tag, err
}

// GetTagByName 根据名称获取标签
func GetTagByName(name string) (Tag, error) {
	var tag Tag
	err := DB.Where("name = ?", name).First(&tag).Error
	return tag, err
}
