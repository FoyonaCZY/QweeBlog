package group

import (
	"errors"
	"github.com/FoyonaCZY/QweeBlog/models"
	"unicode"
)

type UpdateRequest struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Type uint   `json:"type" binding:"required"`
}

type UpdateResponse struct {
	ID uint `json:"id"`
}

// Update 更新用户组
func (req *UpdateRequest) Update() (UpdateResponse, error) {
	//检查用户组是否存在
	group, err := models.GetGroupByID(req.ID)
	if err != nil {
		return UpdateResponse{}, err
	}

	//验证请求
	if !validateGroupUpdateReq(*req) {
		return UpdateResponse{}, errors.New("请求参数错误")
	}

	//更新用户组
	group.Name = req.Name
	group.Type = req.Type
	err = models.DB.Save(&group).Error
	if err != nil {
		return UpdateResponse{}, err
	}

	return UpdateResponse{ID: group.ID}, nil
}

func validateGroupUpdateReq(req UpdateRequest) bool {
	//检查名称是否含有控制字符
	for _, c := range req.Name {
		if unicode.IsControl(c) {
			return false
		}
	}

	//检查类型是否合法
	if req.Type < models.GroupTypeAdmin || req.Type > models.GroupTypeEditor {
		return false
	}

	return true
}
