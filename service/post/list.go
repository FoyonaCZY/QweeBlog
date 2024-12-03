package post

import (
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/pkg/config"
	"github.com/FoyonaCZY/QweeBlog/service/user"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	gorm.Model
	Title      string `json:"title"`
	Content    string `json:"content"`
	User       user.User
	CategoryID uint   `json:"category_id"`
	Tags       []Tag  `json:"tags"`
	Avatar     string `json:"avatar"`
}

type ListResponse struct {
	Count int    `json:"count"`
	Posts []Post `json:"posts"`
}

// List 获取文章列表
func List(pageID int) (ListResponse, error) {
	posts, err := models.GetPosts()
	if err != nil {
		return ListResponse{}, err
	}
	var p []Post
	for i, post := range posts {
		if i < (pageID-1)*config.Configs.Post.PageSize {
			continue
		}
		if i >= pageID*config.Configs.Post.PageSize {
			break
		}
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
			Content: post.Content[:min(config.Configs.Post.SummaryLength, len(post.Content))],
			Avatar:  post.Avatar,
			User: user.User{
				ID:        post.User.ID,
				Nickname:  post.User.Nickname,
				Avatar:    post.User.Avatar,
				GroupType: post.User.Group.Type,
			},
			CategoryID: post.CategoryID,
			Tags:       tags,
		})
	}
	return ListResponse{
		Count: len(p),
		Posts: p,
	}, nil
}

func ListByCategory(categoryID, pageID int) (ListResponse, error) {
	posts, err := models.GetPostsByCategoryID(uint(categoryID))
	if err != nil {
		return ListResponse{}, err
	}
	var p []Post
	for i, post := range posts {
		if i < (pageID-1)*config.Configs.Post.PageSize {
			continue
		}
		if i >= pageID*config.Configs.Post.PageSize {
			break
		}
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
			Content: post.Content[:min(config.Configs.Post.SummaryLength, len(post.Content))],
			Avatar:  post.Avatar,
			User: user.User{
				ID:        post.User.ID,
				Nickname:  post.User.Nickname,
				Avatar:    post.User.Avatar,
				GroupType: post.User.Group.Type,
			},
			CategoryID: post.CategoryID,
			Tags:       tags,
		})
	}
	return ListResponse{
		Count: len(p),
		Posts: p,
	}, nil
}
