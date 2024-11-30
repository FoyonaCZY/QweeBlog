package user

import "github.com/FoyonaCZY/QweeBlog/models"

type ListResponseAdmin struct {
	Count int           `json:"count"`
	Users []models.User `json:"users"`
}

type User struct {
	ID        uint   `json:"id"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	GroupType uint   `json:"group_type"`
}

type ListResponse struct {
	Count int    `json:"count"`
	Users []User `json:"users"`
}

// ListAdmin 管理员获取用户列表
func ListAdmin() (ListResponseAdmin, error) {
	// 获取用户列表
	users, err := models.GetAllUsers()
	if err != nil {
		return ListResponseAdmin{}, err
	}
	return ListResponseAdmin{Count: len(users),
		Users: users}, nil
}

// List 获取用户列表
func List() (ListResponse, error) {
	// 获取用户列表
	users, err := models.GetAllUsers()
	if err != nil {
		return ListResponse{}, err
	}
	var usersResp []User
	for _, user := range users {
		usersResp = append(usersResp, User{
			ID:        user.ID,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			GroupType: user.Group.Type,
		})
	}
	return ListResponse{Count: len(usersResp),
		Users: usersResp}, nil
}
