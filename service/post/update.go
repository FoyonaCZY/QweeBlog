package post

import "github.com/FoyonaCZY/QweeBlog/models"

type UpdateRequest struct {
	ID      uint
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Tags    []Tag  `json:"tags"`
}

type UpdateResponse struct {
	ID uint `json:"id"`
}

// Update 更新文章
func (req *UpdateRequest) Update() (UpdateResponse, error) {
	//找到文章
	post, err := models.GetPostByID(req.ID)
	if err != nil {
		return UpdateResponse{}, err
	}

	//更新文章
	post.Title = req.Title
	post.Content = req.Content
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
