package group

import (
	"github.com/FoyonaCZY/QweeBlog/models"
)

type ListResponse struct {
	Count  int            `json:"count"`
	Groups []models.Group `json:"groups"`
}

func List() (ListResponse, error) {
	// 获取用户列表
	groups, err := models.GetGroups()
	if err != nil {
		return ListResponse{}, err
	}
	return ListResponse{Count: len(groups),
		Groups: groups}, nil
}
