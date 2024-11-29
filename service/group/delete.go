package group

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

// Delete 删除用户组
func (req *DeleteRequest) Delete() (DeleteResponse, error) {
	res := models.DeleteGroupByID(req.ID)
	if res == 0 {
		return DeleteResponse{}, errors.New("找不到用户组")
	}
	return DeleteResponse{ID: req.ID}, nil
}
