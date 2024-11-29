package user

import (
	"errors"
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/util"
	"strings"
	"unicode"
)

type UpdateRequest struct {
	ID           uint   `json:"id" binding:"required"`
	Nickname     string `gorm:"type:varchar(50);not null;" json:"nickname"`
	Email        string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password     string `gorm:"type:varchar(100);not null" json:"password"`
	Avatar       string `gorm:"type:varchar(100);not null" json:"avatar"`
	GroupID      uint   `gorm:"type:int;not null" json:"group_id"`
	ReceiveEmail bool   `gorm:"type:boolean;not null" json:"receive_email"`
}

type UpdateResponse struct {
	ID uint `json:"id"`
}

func (req *UpdateRequest) Update() (UpdateResponse, error) {
	//查找ID是否有效
	user, err := models.GetUserByID(req.ID)
	if err != nil {
		return UpdateResponse{}, errors.New("用户不存在")
	}

	if !ValidateUserUpdateReq(*req) {
		return UpdateResponse{}, errors.New("参数不合法")
	}

	user.Nickname = req.Nickname
	user.Email = req.Email
	user.Avatar = req.Avatar
	user.GroupID = req.GroupID
	user.ReceiveEmail = req.ReceiveEmail
	err = user.SetPassword(req.Password)
	if err != nil {
		return UpdateResponse{}, err
	}
	err = models.DB.Save(&user).Error
	if err != nil {
		return UpdateResponse{}, err
	}
	return UpdateResponse{ID: user.ID}, nil
}

// ValidateUserUpdateReq 验证用户更新请求
func ValidateUserUpdateReq(req UpdateRequest) bool {
	//验证邮箱长度
	if len(req.Email) < 5 || len(req.Email) > 100 {
		return false
	}

	//验证昵称长度
	if len(req.Nickname) < 2 || len(req.Nickname) > 50 {
		return false
	}

	//验证昵称是否含有非法字符
	for _, c := range req.Nickname {
		if unicode.IsControl(c) {
			return false
		}
	}

	//验证密码是否只含有合法字符
	password := req.Password
	for _, c := range password {
		if !strings.ContainsRune(letterRunes, c) {
			return false
		}
	}

	//验证邮箱是否合法
	if !util.IsValidEmail(req.Email) {
		return false
	}

	//验证头像长度
	if len(req.Avatar) < 5 || len(req.Avatar) > 100 {
		return false
	}

	//验证组ID是否有效
	_, err := models.GetGroupByID(req.GroupID)
	if err != nil {
		return false
	}

	//验证密码长度
	return len(password) >= 8 && len(password) <= 20
}
