package comment

import "github.com/FoyonaCZY/QweeBlog/models"

type CreateRequest struct {
	Content string `json:"content" binding:"required"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id" binding:"required"`
}

type CreateResponse struct {
	ID uint `json:"id"`
}

// Create 创建评论
func (req *CreateRequest) Create() (CreateResponse, error) {
	comment := models.Comment{
		Content: req.Content,
		UserID:  req.UserID,
		PostID:  req.PostID,
	}

	if err := models.DB.Create(&comment).Error; err != nil {
		return CreateResponse{}, err
	}

	return CreateResponse{ID: comment.ID}, nil
}
