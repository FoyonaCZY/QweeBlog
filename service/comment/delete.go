package comment

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

// Delete 删除评论
func (req *DeleteRequest) Delete() (DeleteResponse, error) {
	res := models.DeleteCommentByID(req.ID)
	if res == 0 {
		return DeleteResponse{}, errors.New("找不到评论")
	}
	return DeleteResponse{ID: req.ID}, nil
}
