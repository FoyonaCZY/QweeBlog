package category

import (
	"errors"
	"github.com/FoyonaCZY/QweeBlog/models"
	"unicode"
)

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateResponse struct {
	ID uint `json:"id"`
}

// Create 创建分类
func (req *CreateRequest) Create() (CreateResponse, error) {
	if !Validate(req.Name) {
		return CreateResponse{}, errors.New("分类名不合法")
	}

	category := models.Category{
		Name: req.Name,
	}

	if err := models.DB.Create(&category).Error; err != nil {
		return CreateResponse{}, err
	}

	return CreateResponse{ID: category.ID}, nil
}

// Validate 验证分类名合法性
func Validate(name string) bool {
	// 检查分类名长度
	if len(name) < 2 || len(name) > 30 {
		return false
	}

	// 检查分类名是否含有控制字符
	for _, c := range name {
		if unicode.IsControl(c) {
			return false
		}
	}

	return true
}
