package user

import "github.com/FoyonaCZY/QweeBlog/models"

type DeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}
type DeleteResponse struct {
	ID uint `json:"id"`
}

// Delete 删除用户
func (req *DeleteRequest) Delete() (DeleteResponse, error) {
	return DeleteResponse{ID: req.ID}, models.DeleteUserByID(req.ID)
}
