package models

import (
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
