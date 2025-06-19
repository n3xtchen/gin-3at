package dao

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/internal/model"
)

type CategoryDao struct {
	db *gorm.DB
}

// NewCategoryDao creates a new instance of CategoryDao.
func NewCategoryDao(db *gorm.DB) *CategoryDao {
	return &CategoryDao{
		db,
	}
}

// create Category
func (dao *CategoryDao) CreateCategory(category m.Category) error {
	if err := dao.db.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

// list Categories
func (dao *CategoryDao) ListCategories() ([]m.Category, error) {
	var categories []m.Category
	if err := dao.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
