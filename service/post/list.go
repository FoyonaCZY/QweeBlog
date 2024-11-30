package post

import (
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/service/user"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	User    user.User
	Tags    []Tag `json:"tags"`
}

type ListResponse struct {
	Count int    `json:"count"`
	Posts []Post `json:"posts"`
}

// List 获取文章列表
func List() (ListResponse, error) {
	posts, err := models.GetPosts()
	if err != nil {
		return ListResponse{}, err
	}
	var p []Post
	for _, post := range posts {
		var tags []Tag
		for _, tag := range post.Tags {
			tags = append(tags, Tag{
				ID:   tag.ID,
				Name: tag.Name,
			})
		}
		p = append(p, Post{
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
		})
	}
	return ListResponse{
		Count: len(p),
		Posts: p,
	}, nil
}
