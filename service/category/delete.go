package category

import (
	"errors"
	"github.com/FoyonaCZY/QweeBlog/models"
)

type DeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}

type DeleteResponse struct {
	ID uint `json:"id"`
}

// Delete 删除分类
func (req *DeleteRequest) Delete() (DeleteResponse, error) {
	//检查分类下是否有文章
	posts, err := models.GetPostsByCategoryID(req.ID)
	if err != nil {
		return DeleteResponse{}, err
	}
	if len(posts) > 0 {
		return DeleteResponse{}, errors.New("分类下有文章，无法删除")
	}

	//检查是否是唯一的分类
	categories, err := models.GetCategories()
	if err != nil {
		return DeleteResponse{}, err
	}
	if len(categories) == 1 {
		return DeleteResponse{}, errors.New("唯一的分类无法删除")
	}

	RowsAffected := models.DeleteCategoryByID(req.ID)
	if RowsAffected == 0 {
		return DeleteResponse{}, errors.New("找不到此ID的分类")
	}
	return DeleteResponse{ID: req.ID}, nil
}
