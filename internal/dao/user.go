package dao

import (
	"gorm.io/gorm"

	"github.com/n3xtchen/gin-3at/internal/domain/entity"
	"github.com/n3xtchen/gin-3at/internal/domain/repository"
	m "github.com/n3xtchen/gin-3at/internal/model"
)

// UserDao implements the UserRepository interface.
type UserDao struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepositoryImp.
func NewUserDao(db *gorm.DB) repository.UserRepository {
	return &UserDao{
		db: db,
	}
}

// CreateUser creates a new user with the given details.
func (repo *UserDao) Create(user *entity.User) error {

	userModel := m.FromEntityUser(user)

	if err := repo.db.Create(userModel).Error; err != nil {
		return err
	}

	// Update the user entity with the new ID
	user.ID = userModel.ToEntity().ID
	return nil
}

// UpdateUser updates an existing user with the given details.
func (repo *UserDao) Update(user *entity.User, newInfo map[string]interface{}) error {

	userModel := m.FromEntityUser(user)

	if err := repo.db.Model(userModel).Updates(newInfo).Error; err != nil {
		return err
	}
	return nil
}

// ResetPassword resets the user's password.
func (repo *UserDao) ResetPassword(user *entity.User, newPassword string) error {
	userModel := m.FromEntityUser(user)
	userModel.Password = newPassword // Assuming newPassword is already hashed

	if err := repo.db.Model(&userModel).Update("password", userModel.Password).Error; err != nil {
		return err
	}

	// Update the user entity with the new Password
	user.Password = userModel.Password

	return nil
}

// FindBy checks if a user exists by username or email.
func (repo *UserDao) FindBy(usernameOrEmail string) (*entity.User, error) {
	var user entity.User
	if err := repo.db.Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CheckPassword checks if the provided password matches the user's password.
func (repo *UserDao) CheckPassword(user *entity.User, password string) (bool, error) {
	var userModel m.User
	// Assuming user.Password is already hashed
	if err := repo.db.Where("id = ? AND password = ?", user.ID, password).First(&userModel).Error; err != nil {
		return false, err
	}
	return true, nil
}
