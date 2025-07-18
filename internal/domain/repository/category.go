package repository

import (
	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

type CategoryRepository interface {
	// retrieves a category by its ID.
	FindByID(categoryID int) (*e.Category, error)
	// retrieves all categories.
	List() ([]*e.Category, error)
	// creates a new category.
	Save(category *e.Category) error
	// updates an existing category.
	Update(category *e.Category) error
	// deletes a category by its ID.
	Delete(categoryID int) error
}
