package service

import (
	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
	repo "github.com/n3xtchen/gin-3at/internal/domain/repository"
)

// CategoryService defines the methods for category-related operations.
type CategoryService interface {
	// GetCategoryByID retrieves a category by its ID.
	GetCategoryByID(categoryID int) (*e.Category, error)
	// GetAllCategories retrieves all categories.
	GetAllCategories() ([]*e.Category, error)
	// CreateCategory creates a new category.
	CreateCategory(category *e.Category) error
	// UpdateCategory updates an existing category.
	UpdateCategory(category *e.Category) error
	// DeleteCategory deletes a category by its ID.
	DeleteCategory(categoryID int) error
}

// CategoryServiceImp implements the CategoryService interface.
type CategoryServiceImp struct {
	CategoryRepo repo.CategoryRepository
}

// NewCategoryService creates a new instance of CategoryService.
func NewCategoryService(categoryRepo repo.CategoryRepository) CategoryService {
	return &CategoryServiceImp{
		CategoryRepo: categoryRepo,
	}
}

// GetCategoryByID retrieves a category by its ID.
func (s *CategoryServiceImp) GetCategoryByID(categoryID int) (*e.Category, error) {
	return s.CategoryRepo.FindByID(categoryID)
}

// GetAllCategories retrieves all categories.
func (s *CategoryServiceImp) GetAllCategories() ([]*e.Category, error) {
	return s.CategoryRepo.List()
}

// CreateCategory creates a new category.
func (s *CategoryServiceImp) CreateCategory(category *e.Category) error {
	return s.CategoryRepo.Save(category)
}

// UpdateCategory updates an existing category.
func (s *CategoryServiceImp) UpdateCategory(category *e.Category) error {
	return s.CategoryRepo.Update(category)
}

// DeleteCategory deletes a category by its ID.
func (s *CategoryServiceImp) DeleteCategory(categoryID int) error {
	return s.CategoryRepo.Delete(categoryID)
}
