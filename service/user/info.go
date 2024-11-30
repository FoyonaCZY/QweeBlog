package user

import "github.com/FoyonaCZY/QweeBlog/models"

type InfoRequest struct {
	ID uint `json:"id" binding:"required"`
}

type InfoResponseAdmin struct {
	models.User
}

type InfoResponse struct {
	User User
}

// InfoAdmin 管理员获取用户信息
func (req *InfoRequest) InfoAdmin() (InfoResponseAdmin, error) {
	user, err := models.GetUserByID(req.ID)
	if err != nil {
		return InfoResponseAdmin{}, err
	}
	return InfoResponseAdmin{user}, nil
}

// Info 获取用户信息
func (req *InfoRequest) Info() (InfoResponse, error) {
	user, err := models.GetUserByID(req.ID)
	if err != nil {
		return InfoResponse{}, err
	}
	var u User
	u.ID = user.ID
	u.Nickname = user.Nickname
	u.Avatar = user.Avatar
	u.GroupType = user.Group.Type
	return InfoResponse{u}, nil
}
