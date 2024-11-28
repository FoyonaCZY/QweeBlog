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

// GetCommentByID 根据ID获取评论
func GetCommentByID(id uint) (Comment, error) {
	var comment Comment
	err := DB.Where("id = ?", id).First(&comment).Error
	return comment, err
}

// GetCommentsByPostID 根据文章ID获取评论
func GetCommentsByPostID(postID uint) ([]Comment, error) {
	var comments []Comment
	err := DB.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}

// GetCommentsByUserID 根据用户ID获取评论
func GetCommentsByUserID(userID uint) ([]Comment, error) {
	var comments []Comment
	err := DB.Where("user_id = ?", userID).Find(&comments).Error
	return comments, err
}
