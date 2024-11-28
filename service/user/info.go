package user

import "github.com/FoyonaCZY/QweeBlog/models"

type InfoRequest struct {
	ID uint `json:"id" binding:"required"`
}

type InfoResponse struct {
	ID           uint   `json:"id"`
	Nickname     string `json:"nickname"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	GroupID      uint   `json:"group_id"`
	ReceiveEmail bool   `json:"receive_email"`
}

// Info 获取用户信息
func (req *InfoRequest) Info() (InfoResponse, error) {
	user, err := models.GetUserByID(req.ID)
	if err != nil {
		return InfoResponse{}, err
	}
	return InfoResponse{
		ID:           user.ID,
		Nickname:     user.Nickname,
		Email:        user.Email,
		Avatar:       user.Avatar,
		GroupID:      user.GroupID,
		ReceiveEmail: user.ReceiveEmail,
	}, nil
}
