package group

import "github.com/FoyonaCZY/QweeBlog/models"

type InfoRequest struct {
	ID uint `json:"id" binding:"required"`
}

type InfoResponse struct {
	models.Group
}

// Info 获取用户组信息
func (req *InfoRequest) Info() (InfoResponse, error) {
	group, err := models.GetGroupByID(req.ID)
	if err != nil {
		return InfoResponse{}, err
	}
	return InfoResponse{group}, nil
}
