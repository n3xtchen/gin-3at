package dao

import (
	m "github.com/n3xtchen/gin-3at/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao() *AddressDao {
	return &AddressDao{
		DB: DB,
	}
}

// CreateAddress creates a new address in the database.
func (dao *AddressDao) CreateAddress(address m.Address) error {
	if err := dao.DB.Create(&address).Error; err != nil {
		return err
	}
	return nil
}

// GetOrCreateAddressID retrieves the address ID if it exists, or creates a new address and returns its ID.
func (dao *AddressDao) GetOrCreateAddressID(address m.Address) (uint, error) {
	var existingAddress m.Address
	// Check if the address already exists in the database.
	if err := DB.Where("user_id = ? AND name = ? AND phone = ? AND address = ?", address.UserID, address.Name, address.Phone, address.Address).First(&existingAddress).Error; err == nil {
		// Address exists, return its ID.
		return existingAddress.ID, nil
	} else if err.Error() != "record not found" {
		// If the error is not "record not found", return the error.
		return 0, err
	}

	// Address does not exist, create a new one.
	if err := DB.Create(&address).Error; err != nil {
		return 0, err
	}

	return address.ID, nil
}
