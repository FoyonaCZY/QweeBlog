package group

import (
	"github.com/FoyonaCZY/QweeBlog/models"
)

type ListResponse struct {
	Groups []models.Group `json:"groups"`
}

func List() (ListResponse, error) {
	// 获取用户列表
	groups, err := models.GetGroups()
	if err != nil {
		return ListResponse{}, err
	}
	return ListResponse{Groups: groups}, nil
}
