package user

import (
	"errors"
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/pkg/mail"
	"github.com/FoyonaCZY/QweeBlog/util"
	"strings"
	"unicode"
)

const letterRunes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+"

type RegisterRequest struct {
	Nickname     string `json:"nickname" binding:"required,min=2,max=50"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6,max=20"`
	ReceiveEmail bool   `json:"receive_email"`
}

type RegisterResponse struct {
	ID uint `json:"id"`
}

// Register 用户注册
func (req *RegisterRequest) Register() (RegisterResponse, error) {
	if !ValidateUserRegisterReq(*req) {
		return RegisterResponse{}, errors.New("参数错误")
	}

	user := models.NewDefaultUser()
	user.Nickname = req.Nickname
	user.Email = req.Email
	if err := user.SetPassword(req.Password); err != nil {
		return RegisterResponse{}, errors.New("密码哈希失败")
	}
	user.ReceiveEmail = req.ReceiveEmail
	err := models.DB.Create(&user).Error
	if err != nil {
		return RegisterResponse{}, err
	}

	// 发送激活邮件
	go func() {
		_ = mail.SendActivationEmail(user.Email)
	}()

	return RegisterResponse{ID: user.ID}, nil
}

// ValidateUserRegisterReq 验证用户注册请求
func ValidateUserRegisterReq(req RegisterRequest) bool {
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

	//验证邮箱是否合法
	if !util.IsValidEmail(req.Email) {
		return false
	}

	//验证密码是否只含有合法字符
	password := req.Password
	for _, c := range password {
		if !strings.ContainsRune(letterRunes, c) {
			return false
		}
	}

	//验证密码长度
	return len(password) >= 8 && len(password) <= 20
}
