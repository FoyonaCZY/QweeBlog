package group

import (
	"errors"
	"github.com/FoyonaCZY/QweeBlog/models"
	"unicode"
)

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Type uint   `json:"type" binding:"required"`
}

type CreateResponse struct {
	ID uint `json:"id"`
}

// Create 创建用户组
func (req *CreateRequest) Create() (CreateResponse, error) {
	if !ValidateGroupCreateReq(req) {
		return CreateResponse{}, errors.New("参数不合法")
	}

	group := models.Group{
		Name: req.Name,
		Type: req.Type,
	}
	if err := models.DB.Create(&group).Error; err != nil {
		return CreateResponse{}, err
	}
	return CreateResponse{ID: group.ID}, nil
}

func ValidateGroupCreateReq(req *CreateRequest) bool {
	// 检查用户组类型
	if req.Type < models.GroupTypeAdmin || req.Type > models.GroupTypeEditor {
		return false
	}

	// 检查用户组名称长度
	if len(req.Name) < 2 || len(req.Name) > 50 {
		return false
	}

	//检查用户组名称是否含有控制字符
	for _, c := range req.Name {
		if unicode.IsControl(c) {
			return false
		}
	}

	return true
}
