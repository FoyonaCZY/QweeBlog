package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

// GetCategoryByID 根据ID获取分类
func GetCategoryByID(id uint) (Category, error) {
	var category Category
	if err := DB.First(&category, id).Error; err != nil {
		return Category{}, err
	}
	return category, nil
}

// GetCategories 获取分类列表
func GetCategories() ([]Category, error) {
	var categories []Category
	if err := DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// DeleteCategoryByID 根据ID删除分类
func DeleteCategoryByID(id uint) int64 {
	return DB.Delete(Category{}, id).RowsAffected
}
