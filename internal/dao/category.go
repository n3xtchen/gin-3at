package dao

import (
	"gorm.io/gorm"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
	repo "github.com/n3xtchen/gin-3at/internal/domain/repository"
	m "github.com/n3xtchen/gin-3at/internal/model"
)

type CategoryDao struct {
	db *gorm.DB
}

// NewCategoryDao creates a new instance of CategoryDao.
func NewCategoryDao(db *gorm.DB) repo.CategoryRepository {
	return &CategoryDao{
		db,
	}
}

// create Category
func (dao *CategoryDao) Save(category *e.Category) error {
	// Convert the entity to Model
	categoryModel := m.FromEntityCategory(category)
	if err := dao.db.Create(&categoryModel).Error; err != nil {
		return err
	}
	return nil
}

// list Categories
func (dao *CategoryDao) List() ([]*e.Category, error) {
	var categories []m.Category
	if err := dao.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	// Convert the model to entity
	var entities []*e.Category
	for _, category := range categories {
		categoryEntity := category.ToEntity()
		entities = append(entities, &categoryEntity)
	}
	return entities, nil
}

// FindByID retrieves a category by its ID.
func (dao *CategoryDao) FindByID(categoryID int) (*e.Category, error) {
	var category m.Category
	if err := dao.db.First(&category, categoryID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	// Convert the model to entity
	categoryEntity := category.ToEntity()

	return &categoryEntity, nil
}

// Update modifies an existing category.
func (dao *CategoryDao) Update(category *e.Category) error {
	// Convert the entity to Model
	categoryModel := m.FromEntityCategory(category)
	if err := dao.db.Save(&categoryModel).Error; err != nil {
		return err
	}
	return nil
}

// Delete removes a category by its ID.
func (dao *CategoryDao) Delete(categoryID int) error {
	if err := dao.db.Delete(&m.Category{}, categoryID).Error; err != nil {
		return err
	}
	return nil
}
