package category

import "github.com/FoyonaCZY/QweeBlog/models"

type ListResponse struct {
	Count      int               `json:"count"`
	Categories []models.Category `json:"categories"`
}

// List 获取分类列表
func List() (ListResponse, error) {
	// 获取分类列表
	categories, err := models.GetCategories()
	if err != nil {
		return ListResponse{}, err
	}
	return ListResponse{Count: len(categories),
		Categories: categories}, nil
}
