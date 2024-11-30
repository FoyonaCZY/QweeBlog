package category

import (
	"errors"
	"github.com/FoyonaCZY/QweeBlog/models"
)

type UpdateRequest struct {
	ID   uint
	Name string `json:"name" binding:"required"`
}

type UpdateResponse struct {
	ID uint `json:"id"`
}

// Update 更新分类
func (req *UpdateRequest) Update() (UpdateResponse, error) {
	if !Validate(req.Name) {
		return UpdateResponse{}, errors.New("参数不合法")
	}

	category, err := models.GetCategoryByID(req.ID)
	if err != nil {
		return UpdateResponse{}, err
	}

	category.Name = req.Name
	if err := models.DB.Save(&category).Error; err != nil {
		return UpdateResponse{}, err
	}
	return UpdateResponse{ID: category.ID}, nil
}
