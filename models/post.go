package models

import (
	"github.com/jinzhu/gorm"
)

// Post 文章模型
type Post struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null;" json:"title"`
	Content     string `gorm:"type:text;not null;" json:"content"`
	UserID      uint   `gorm:"type:int;not null" json:"user_id"`
	IsPublished bool   `gorm:"type:boolean;not null" json:"is_published"`

	// 关联模型
	User User  `gorm:"save_associations:false:false"`
	Tag  []Tag `gorm:"many2many:post_tags;save_associations:false:false"`
}

// GetPostByID 根据ID获取文章
func GetPostByID(id uint) (Post, error) {
	var post Post
	err := DB.Where("id = ?", id).First(&post).Error
	return post, err
}

// GetPostsByUserID 根据用户ID获取文章
func GetPostsByUserID(userID uint) ([]Post, error) {
	var posts []Post
	err := DB.Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

// GetPostsByTagID 根据标签ID获取文章
func GetPostsByTagID(tagID uint) ([]Post, error) {
	var posts []Post
	err := DB.Where("tag_id = ?", tagID).Find(&posts).Error
	return posts, err
}

// GetPostsByTagNames 根据标签名获取文章
func GetPostsByTagNames(tagNames []string) ([]Post, error) {
	var posts []Post
	err := DB.Where("tag_name IN (?)", tagNames).Find(&posts).Error
	return posts, err
}

// GetPostsByTagIDs 根据标签ID获取包含所有标签的文章
func GetPostsByTagIDs(tagIDs []uint) ([]Post, error) {
	var posts []Post
	// 使用 JOIN 查询来查找包含所有标签的文章
	err := DB.Table("posts").
		Joins("JOIN post_tags ON posts.id = post_tags.post_id").
		Where("post_tags.tag_id IN (?)", tagIDs).
		Group("posts.id").
		Having("COUNT(DISTINCT post_tags.tag_id) = ?", len(tagIDs)).
		Find(&posts).Error
	return posts, err
}
