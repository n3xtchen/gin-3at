package dto

import (
	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

type CategoryReq struct {
	CategoryID int `uri:"id" binding:"required"` // ID of the CategoryID
}

type CategroyRes struct {
	CategoryID  int    `json:"category_id"` // ID of the CategoryID
	Name        string `json:"name"`        // Name of the CategoryID
	Description string `json:"description"` // Description of the CategoryID
}

func ToCategoryRes(category *e.Category) CategroyRes {
	return CategroyRes{
		CategoryID:  category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}
