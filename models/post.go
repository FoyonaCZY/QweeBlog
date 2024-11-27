package models

import (
	"github.com/jinzhu/gorm"
)

// Post 文章模型
type Post struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null;" json:"title"`
	Content string `gorm:"type:text;not null;" json:"content"`
	UserID  uint   `gorm:"type:int;not null" json:"user_id"`

	// 关联模型
	User User `gorm:"save_associations:false:false"`
}
