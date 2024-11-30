package post

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

// Delete 删除文章
func (req *DeleteRequest) Delete() (DeleteResponse, error) {
	RowsAffected := models.DeletePostByID(req.ID)
	if RowsAffected == 0 {
		return DeleteResponse{}, errors.New("找不到此ID的文章")
	}
	return DeleteResponse{ID: req.ID}, nil
}
