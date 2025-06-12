package dao

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/model"
)

type CategoryDao struct {
	*gorm.DB
}

// NewCategoryDao creates a new instance of CategoryDao.
func NewCategoryDao() *CategoryDao {
	return &CategoryDao{
		DB: DB,
	}
}

// create Category
func (dao *CategoryDao) CreateCategory(category m.Category) error {
	if err := dao.DB.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

// list Categories
func (dao *CategoryDao) ListCategories() ([]m.Category, error) {
	var categories []m.Category
	if err := dao.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
