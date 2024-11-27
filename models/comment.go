package models

import (
	"github.com/jinzhu/gorm"
)

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null;" json:"content"`
	UserID  uint   `gorm:"type:int;not null" json:"user_id"`
	PostID  uint   `gorm:"type:int;not null" json:"post_id"`

	// 关联模型
	User User `gorm:"save_associations:false:false"`
}
