package user

import "github.com/FoyonaCZY/QweeBlog/models"

type InfoRequest struct {
	ID uint `json:"id" binding:"required"`
}

type InfoResponse struct {
	models.User
}

// Info 获取用户信息
func (req *InfoRequest) Info() (InfoResponse, error) {
	user, err := models.GetUserByID(req.ID)
	if err != nil {
		return InfoResponse{}, err
	}
	return InfoResponse{user}, nil
}
