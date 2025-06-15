package dao

import (
	"gorm.io/gorm"

	m "github.com/n3xtchen/gin-3at/internal/model"
)

type AddressDao struct {
	db *gorm.DB
}

func NewAddressDao(db *gorm.DB) *AddressDao {
	return &AddressDao{
		db,
	}
}

// CreateAddress creates a new address in the database.
func (dao *AddressDao) CreateAddress(address m.Address) error {
	if err := dao.db.Create(&address).Error; err != nil {
		return err
	}
	return nil
}

// GetOrCreateAddressID retrieves the address ID if it exists, or creates a new address and returns its ID.
func (dao *AddressDao) GetOrCreateAddressID(address m.Address) (uint, error) {
	var existingAddress m.Address
	// Check if the address already exists in the database.
	if err := dao.db.Where("user_id = ? AND name = ? AND phone = ? AND address = ?", address.UserID, address.Name, address.Phone, address.Address).First(&existingAddress).Error; err == nil {
		// Address exists, return its ID.
		return existingAddress.ID, nil
	}

	// Address does not exist, create a new one.
	if err := dao.db.Create(&address).Error; err != nil {
		return 0, err
	}

	return address.ID, nil
}
