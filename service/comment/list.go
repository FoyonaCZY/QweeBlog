package comment

import (
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/service/user"
)

type Comment struct {
	ID      uint      `json:"id"`
	Content string    `json:"content"`
	PostID  uint      `json:"post_id"`
	User    user.User `json:"user"`
}

type ListResponse struct {
	Count    int       `json:"count"`
	Comments []Comment `json:"comments"`
}

// List 获取评论列表
func List() (ListResponse, error) {
	// 获取评论列表
	comments, err := models.GetComments()
	if err != nil {
		return ListResponse{}, err
	}

	var commentsResp []Comment
	for _, comment := range comments {
		commentsResp = append(commentsResp, Comment{
			ID:      comment.ID,
			Content: comment.Content,
			PostID:  comment.PostID,
			User: user.User{
				ID:       comment.User.ID,
				Nickname: comment.User.Nickname,
				Avatar:   comment.User.Avatar,
			},
		})
	}
	return ListResponse{Count: len(commentsResp),
		Comments: commentsResp}, nil
}

// ListByPostID 根据文章ID获取评论列表
func ListByPostID(postID uint) (ListResponse, error) {
	// 获取评论列表
	comments, err := models.GetCommentsByPostID(postID)
	if err != nil {
		return ListResponse{}, err
	}

	var commentsResp []Comment
	for _, comment := range comments {
		commentsResp = append(commentsResp, Comment{
			ID:      comment.ID,
			Content: comment.Content,
			PostID:  comment.PostID,
			User: user.User{
				ID:       comment.User.ID,
				Nickname: comment.User.Nickname,
				Avatar:   comment.User.Avatar,
			},
		})
	}
	return ListResponse{Count: len(commentsResp),
		Comments: commentsResp}, nil
}
