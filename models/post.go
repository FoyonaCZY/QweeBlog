package models

import (
	"github.com/FoyonaCZY/QweeBlog/pkg/config"
	"github.com/jinzhu/gorm"
)

// Post 文章模型
type Post struct {
	gorm.Model
	Title      string `gorm:"type:varchar(100);not null;" json:"title"`
	Content    string `gorm:"type:text;not null;" json:"content"`
	UserID     uint   `gorm:"type:int;not null" json:"user_id"`
	CategoryID uint   `gorm:"type:int;not null" json:"category_id"`
	Avatar     string `gorm:"type:varchar(255);not null" json:"avatar"`

	// 关联模型
	User User  `gorm:"save_associations:false:false"`
	Tags []Tag `gorm:"many2many:post_tags;save_associations:false:false"`
}

// GetPostByID 根据ID获取文章
func GetPostByID(id uint) (Post, error) {
	var post Post
	err := DB.Where("id = ?", id).Preload("User").Preload("Tags").First(&post).Error
	return post, err
}

// GetPostsByUserID 根据用户ID获取文章
func GetPostsByUserID(userID uint) ([]Post, error) {
	var posts []Post
	err := DB.Where("user_id = ?", userID).Preload("User").Preload("Tags").Find(&posts).Error
	return posts, err
}

// GetPosts 获取文章列表
func GetPosts() ([]Post, error) {
	var posts []Post
	err := DB.Preload("User").Preload("Tags").Find(&posts).Error
	return posts, err
}

// DeletePostByID 根据ID删除文章
func DeletePostByID(ID uint) int64 {
	return DB.Delete(Post{}, ID).RowsAffected
}

// GetPostsByCategoryID 根据分类ID获取文章
func GetPostsByCategoryID(categoryID uint) ([]Post, error) {
	var posts []Post
	err := DB.Where("category_id = ?", categoryID).Preload("User").Preload("Tags").Find(&posts).Error
	return posts, err
}

// NewPost 新建文章
func NewPost() Post {
	return Post{
		Avatar: config.Configs.DefaultAvatar.DefaultPostAvatar,
	}
}
