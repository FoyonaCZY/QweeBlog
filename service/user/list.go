package user

import "github.com/FoyonaCZY/QweeBlog/models"

type ListResponse struct {
	Users []models.User `json:"users"`
}

func List() (ListResponse, error) {
	// 获取用户列表
	users, err := models.GetAllUsers()
	if err != nil {
		return ListResponse{}, err
	}
	return ListResponse{Users: users}, nil
}
