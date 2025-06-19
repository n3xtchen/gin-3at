package dao

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/internal/model"
)

type ProductDao struct {
	db *gorm.DB
}

func NewProductDao(db *gorm.DB) *ProductDao {
	return &ProductDao{
		db,
	}
}

// CreateProduct creates a new product in the database.
func (dao *ProductDao) CreateProduct(product m.Product) error {
	if err := dao.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

// GetProducts retrieves all products from the database.
// by pagination, you can add limit and offset parameters.
// with Category, you can use Preload to load the category information.
func (dao *ProductDao) GetProducts(limit, offset int) ([]m.Product, error) {
	var products []m.Product
	if err := dao.db.Preload("Category").Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// GetProductByID retrieves a product by its ID.
// with Category, you can use Preload to load the category information.
func (dao *ProductDao) GetProductByID(productID uint) (*m.Product, error) {
	var product m.Product
	if err := dao.db.Preload("Category").First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// GetProductByCategoryID retrieves products by category ID.
// by pagination, you can add limit and offset parameters.
// with Category, you can use Preload to load the category information.
func (dao *ProductDao) GetProductByCategoryID(categoryID uint, limit, offset int) ([]m.Product, error) {
	var products []m.Product
	if err := dao.db.Where("category_id = ?", categoryID).Preload("Category").Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
