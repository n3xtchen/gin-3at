package service

import (
	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
	repo "github.com/n3xtchen/gin-3at/internal/domain/repository"
)

var (
	// ErrUserNotFound is returned when a user is not found in the system.
	ErrUserNotFound = &BizError{Code: 10001, Message: "User not found"}
	// ErrInvalidPassword is returned when the provided password does not match.
	ErrInvalidPassword = &BizError{Code: 10002, Message: "Invalid password"}
	// ErrUserAlreadyExists is returned when trying to register a user that already exists.
	ErrUserAlreadyExists = &BizError{Code: 10003, Message: "User already exists"}
	// ErrUserRegistrationFailed is returned when user registration fails.
	ErrUserRegistrationFailed = &BizError{Code: 10004, Message: "User registration failed"}
	// ErrUserResetPasswordFailed is returned when resetting a user's password fails.
	ErrUserResetPasswordFailed = &BizError{Code: 10005, Message: "User reset password failed"}
)

type UserService interface {

	// Register a new User
	Register(user *e.User) error

	// Reset Password
	ResetPassword(username, newPassword string) error

	// Login
	Login(username, password string) (*e.User, error)

	// Logout
	Logout(userID int) error
}

type UserServiceImp struct {
	UserRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return &UserServiceImp{
		UserRepo: userRepo,
	}
}

func (s *UserServiceImp) Register(user *e.User) error {

	// Check if the user already existingUser
	_, err := s.UserRepo.FindBy(user.Username)
	if err != nil {
		return ErrUserRegistrationFailed
	}

	err = s.UserRepo.Create(user)
	if err != nil {
		return ErrUserRegistrationFailed
	}

	// Check if the user already existing
	_, err = s.UserRepo.FindBy(user.Username)
	if err != nil {
		return ErrUserRegistrationFailed
	}

	return nil
}

func (s *UserServiceImp) ResetPassword(username, newPassword string) error {
	// Check if the user exists
	existingUser, err := s.UserRepo.FindBy(username)
	if err != nil || existingUser == nil {
		return ErrUserNotFound
	}

	// Reset the password
	err = s.UserRepo.ResetPassword(existingUser, newPassword)
	if err != nil {
		return ErrUserResetPasswordFailed
	}

	return nil
}

func (s *UserServiceImp) Login(username, password string) (*e.User, error) {

	// Check if the use exists
	existingUser, err := s.UserRepo.FindBy(username)
	if err != nil || existingUser == nil {
		return nil, ErrUserNotFound
	}

	// Check if the password matches
	passwordMatch, err := s.UserRepo.CheckPassword(existingUser, password)
	if err != nil || !passwordMatch {
		return nil, ErrInvalidPassword
	}

	return existingUser, nil
}

func (s *UserServiceImp) Logout(userID int) error {
	// Implement logout logic if needed, e.g., clearing session or token
	// For now, we can just return nil as there's no specific action required
	return nil
}
