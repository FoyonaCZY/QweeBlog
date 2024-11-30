package post

import (
	"errors"
	"github.com/FoyonaCZY/QweeBlog/models"
	"unicode"
)

type CreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint   `json:"user_id" binding:"required"`
	Tags    []uint `json:"tags"`
}

type CreateResponse struct {
	ID uint `json:"id"`
}

// Create 创建文章
func (req *CreateRequest) Create() (CreateResponse, error) {
	if !ValidatePostCreateReq(req) {
		return CreateResponse{}, errors.New("参数不合法")
	}

	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}
	for _, tagID := range req.Tags {
		tag, err := models.GetTagByID(tagID)
		if err != nil {
			return CreateResponse{}, err
		}
		post.Tags = append(post.Tags, tag)
	}

	if err := models.DB.Create(&post).Error; err != nil {
		return CreateResponse{}, err
	}

	return CreateResponse{ID: post.ID}, nil
}

func ValidatePostCreateReq(req *CreateRequest) bool {
	// 检查标题长度
	if len(req.Title) < 2 || len(req.Title) > 100 {
		return false
	}
	//检查标题是否含有控制字符
	for _, c := range req.Title {
		if unicode.IsControl(c) {
			return false
		}
	}

	//检查标签是否存在
	for _, tagID := range req.Tags {
		tag, err := models.GetTagByID(tagID)
		if err != nil {
			return false
		}
		if tag.ID == 0 {
			return false
		}
	}

	return true
}
