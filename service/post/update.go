package post

import (
	"errors"
	"github.com/FoyonaCZY/QweeBlog/models"
)

type UpdateRequest struct {
	ID         uint
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
	Avatar     string `json:"avatar"`
	Tags       []Tag  `json:"tags"`
}

type UpdateResponse struct {
	ID uint `json:"id"`
}

// Update 更新文章
func (req *UpdateRequest) Update() (UpdateResponse, error) {
	//检查内容合法性
	r := &CreateRequest{
		Title:      req.Title,
		Content:    req.Content,
		CategoryID: req.CategoryID,
	}
	for _, tag := range req.Tags {
		r.Tags = append(r.Tags, tag.ID)
	}
	if !ValidatePostCreateReq(r) {
		return UpdateResponse{}, errors.New("参数不合法")
	}

	//找到文章
	post, err := models.GetPostByID(req.ID)
	if err != nil {
		return UpdateResponse{}, err
	}

	//更新文章
	post.Title = req.Title
	post.Content = req.Content
	post.Avatar = req.Avatar
	post.CategoryID = req.CategoryID
	for _, tag := range req.Tags {
		t, err := models.GetTagByID(tag.ID)
		if err != nil {
			return UpdateResponse{}, err
		}
		post.Tags = append(post.Tags, t)
	}
	if err := models.DB.Save(&post).Error; err != nil {
		return UpdateResponse{}, err
	}
	return UpdateResponse{ID: post.ID}, nil
}
