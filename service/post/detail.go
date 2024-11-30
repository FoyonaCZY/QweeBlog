package post

import (
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/service/user"
)

type DetailRequest struct {
	ID uint `json:"id" binding:"required"`
}

type DetailResponse struct {
	Post
}

// Detail 获取文章详情
func (req *DetailRequest) Detail() (DetailResponse, error) {
	post, err := models.GetPostByID(req.ID)
	if err != nil {
		return DetailResponse{}, err
	}
	var p Post
	var tags []Tag
	for _, tag := range post.Tags {
		tags = append(tags, Tag{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}
	p = Post{
		Model:   post.Model,
		Title:   post.Title,
		Content: post.Content,
		User: user.User{
			ID:        post.User.ID,
			Nickname:  post.User.Nickname,
			Avatar:    post.User.Avatar,
			GroupType: post.User.Group.Type,
		},
		Tags: tags,
	}
	return DetailResponse{p}, nil
}
