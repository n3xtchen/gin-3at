package dao

import (
	"testing"

	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
	m "github.com/n3xtchen/gin-3at/internal/model"
)

// Test Create User
func TestUserDaoCreateUser(t *testing.T) {
	dao := NewUserDao(db)

	// Create a new user
	user := &e.User{
		Username: "testuser",
		Password: "testpassword",
		Email:    "testUser@email.com",
	}
	err := dao.Create(user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Verify the user was created
	var createdUser m.User

	err = db.Where("username = ?", user.Username).First(&createdUser).Error
	if err != nil {
		t.Fatalf("Failed to find created user: %v", err)
	}
	if createdUser.Username != user.Username || createdUser.Email != user.Email {
		t.Fatalf("Created user does not match expected values: got %v, want %v", createdUser, user)
	}
	// Log success
	t.Logf("User %s created successfully with ID %d", user.Username, createdUser.ID)
}

// Test Find User By Username or Email
func TestUserDaoFindBy(t *testing.T) {
	dao := NewUserDao(db)

	// Create a new user
	user := &e.User{
		Username: "finduser",
		Password: "findpassword",
		Email:    "finduser@mail.com",
	}

	err := dao.Create(user)
	if err != nil {
		t.Fatalf("Failed to create user for find test: %v", err)
	}
	// Find the user by Username
	foundUser, err := dao.FindBy(user.Username)
	if err != nil {
		t.Fatalf("Failed to find user by username: %v", err)
	}
	if foundUser.Username != user.Username || foundUser.Email != user.Email {
		t.Fatalf("Found user does not match expected values: got %v, want %v", foundUser, user)
	}
	// Find the user by Email
	foundUserByEmail, err := dao.FindBy(user.Email)
	if err != nil {
		t.Fatalf("Failed to find user by email: %v", err)
	}
	if foundUserByEmail.Username != user.Username || foundUserByEmail.Email != user.Email {
		t.Fatalf("Found user by email does not match expected values: got %v, want %v", foundUserByEmail, user)
	}
	// Log successfully
	t.Logf("User %s found successfully with ID %d", user.Username, foundUser.ID)
}

// Test Check Password
func TestUserDaoCheckPassword(t *testing.T) {
	dao := NewUserDao(db)

	// Create a new user
	user := &e.User{
		Username: "checkuser",
		Password: "checkpassword",
		Email:    "checkuser@mail.com",
	}

	err := dao.Create(user)
	if err != nil {
		t.Fatalf("Failed to create user for password check test: %v", err)
	}
	// Check the Password
	isValid, err := dao.CheckPassword(user, user.Password)
	if err != nil {
		t.Fatalf("Failed to check password: %v", err)
	}
	if !isValid {
		t.Fatalf("Password check failed for user %s", user.Username)
	}
	// Log successfully
	t.Logf("Password for user %s is valid", user.Username)
}

// Test Reset Password
func TestUserDaoResetPassword(t *testing.T) {
	dao := NewUserDao(db)

	// Create a new user
	user := &e.User{
		Username: "resetuser",
		Password: "oldpassword",
		Email:    "resetuser@mail.com",
	}

	err := dao.Create(user)
	if err != nil {
		t.Fatalf("Failed to create user for password reset test: %v", err)
	}
	// Reset the Password
	newPassword := "newpassword"

	err = dao.ResetPassword(user, newPassword)
	if err != nil {
		t.Fatalf("Failed to reset password: %v", err)
	}

	// Verify the Password was Reset
	isValid, err := dao.CheckPassword(user, newPassword)
	if err != nil {
		t.Fatalf("Failed to check password after reset: %v", err)
	}

	if !isValid {
		t.Fatalf("Password reset failed for user %s", user.Username)
	}

	// Log successfully
	t.Logf("Password for user %s reset successfully", user.Username)
}

// Test Update User Or Error
func TestUserDaoUpdateUser(t *testing.T) {
	dao := NewUserDao(db)

	// Create a new user
	user := &e.User{
		Username: "updateuser",
		Password: "updatepassword",
		Email:    "updateuser@mail.com",
	}

	err := dao.Create(user)
	if err != nil {
		t.Fatalf("Failed to create user for update test: %v", err)
	}

	// Update the User
	newInfo := map[string]interface{}{
		"Email": "new updateuser@mail.com",
	}

	err = dao.Update(user, newInfo)
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}
	// Verify the User was Updated
	var updatedUser m.User
	err = db.Where("username = ?", user.Username).First(&updatedUser).Error
	if err != nil {
		t.Fatalf("Failed to find updated user: %v", err)
	}
	if updatedUser.Email != newInfo["Email"] {
		t.Fatalf("Updated user does not match expected values: got %v, want %v", updatedUser.Email, newInfo["Email"])
	}
	// Log successfully
	t.Logf("User %s updated successfully with new email %s", user.Username, updatedUser.Email)

}
