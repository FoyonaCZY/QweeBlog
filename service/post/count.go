package post

import (
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/pkg/config"
)

type CountResponse struct {
	Count int `json:"count"`
}

// Count 获取文章数量
func Count() (CountResponse, error) {
	posts, err := models.GetPosts()
	count := len(posts) / config.Configs.Post.PageSize
	if len(posts)%config.Configs.Post.PageSize != 0 {
		count++
	}
	if err != nil {
		return CountResponse{}, err
	}
	return CountResponse{Count: count}, nil
}