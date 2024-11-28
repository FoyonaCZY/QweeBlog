package user

import (
	"errors"
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/pkg/auth"
	"github.com/FoyonaCZY/QweeBlog/util"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4,max=64"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// Login 用户登录
func (req *LoginRequest) Login() (LoginResponse, error) {
	// 检查邮箱密码
	if err := CheckEmailAndPassword(req.Email, req.Password); err != nil {
		return LoginResponse{}, errors.New("邮箱或密码错误")
	}

	// 获取用户
	user, err := models.GetUserByEmail(req.Email)
	if err != nil {
		return LoginResponse{}, err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return LoginResponse{}, err
	}
	return LoginResponse{Token: token, User: user}, nil
}

// CheckEmailAndPassword 检查邮箱密码
func CheckEmailAndPassword(email, password string) error {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return err
	}

	return util.ComparePassword(user.Password, password)
}
