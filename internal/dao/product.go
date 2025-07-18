package dao

import (
	"gorm.io/gorm"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
	repo "github.com/n3xtchen/gin-3at/internal/domain/repository"
	m "github.com/n3xtchen/gin-3at/internal/model"
)

type ProductDao struct {
	db *gorm.DB
}

func NewProductDao(db *gorm.DB) repo.ProductRepository {
	return &ProductDao{
		db,
	}
}

// CreateProduct creates a new product in the database.
func (dao *ProductDao) Save(product *e.Product) error {
	// Convert e.Product to m.product
	productModel := m.FromEntityProduct(product)
	if err := dao.db.Create(&productModel).Error; err != nil {
		return err
	}
	return nil
}

// UpdateProduct updates an existing product in the database.
func (dao *ProductDao) Update(product *e.Product) error {
	// Convert e.Product to m.product
	productModel := m.FromEntityProduct(product)
	if err := dao.db.Save(&productModel).Error; err != nil {
		return err
	}
	// Update the product entity with the new values
	*product = productModel.ToEntity()
	// Log the successful Update
	return nil
}

// DeleteProduct deletes a product by its ID from the database.
func (dao *ProductDao) Delete(productID int) error {
	var productModel m.Product
	if err := dao.db.First(&productModel, productID).Error; err != nil {
		return err
	}
	if err := dao.db.Delete(&productModel).Error; err != nil {
		return err
	}
	// Log the successful deletion
	// log.Printf("Product deleted successfully with ID: %d", productID)
	return nil
}

// GetProducts retrieves all products from the database.
// by pagination, you can add limit and offset parameters.
// with Category, you can use Preload to load the category information.
func (dao *ProductDao) List() ([]*e.Product, error) {
	var products []m.Product
	if err := dao.db.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}

	// Convert m.Product to e.Product
	var productEntities []*e.Product
	for _, product := range products {
		// Convert m.Product to e.Product
		productEntity := product.ToEntity()
		productEntities = append(productEntities, &productEntity)
	}
	// if len(productEntities) == 0 {
	// 	return nil, gorm.ErrRecordNotFound
	// }

	return productEntities, nil
}

// GetProductByID retrieves a product by its ID.
// with Category, you can use Preload to load the category information.
func (dao *ProductDao) FindByID(productID int) (*e.Product, error) {
	var productModel m.Product
	if err := dao.db.Preload("Category").First(&productModel, productID).Error; err != nil {
		return nil, err
	}
	// Convert m.Product to e.Product
	product := productModel.ToEntity()
	return &product, nil
}

// GetProductByCategoryID retrieves products by category ID.
// by pagination, you can add limit and offset parameters.
// with Category, you can use Preload to load the category information.
func (dao *ProductDao) ListByCategoryID(categoryID uint, limit, offset int) ([]m.Product, error) {
	var products []m.Product
	if err := dao.db.Where("category_id = ?", categoryID).Preload("Category").Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
