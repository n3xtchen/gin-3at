package dto

import (
	e "github.com/n3xtchen/gin-3at/internal/domain/entity"
)

// RegisterUserRequest represents the request payload for user registration.
type RegisterUserReq struct {
	Username string `json:"username" binding:"required,min=3,max=20"` // Username of the User
	Password string `json:"password" binding:"required,min=6,max=20"` // Password of the User
	Email    string `json:"email" binding:"required,email"`           // Email of the User
}

func (r *RegisterUserReq) ToEntity() *e.User {
	return &e.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}

// LoginUserRequest represents the request payload for user login.
type LoginUserReq struct {
	Username string `json:"username" binding:"required,min=3,max=20"` // Username of the User
	Password string `json:"password" binding:"required,min=6,max=20"` // Password of the User
}

// ResetPasswordRequest represents the request payload for resetting a user's password.
type ResetPasswordReq struct {
	Username    string `json:"username" binding:"required,min=3,max=20"`     // Username of the User
	OldPassword string `json:"old_password" binding:"required,min=6,max=20"` // Old Password of the User
	NewPassword string `json:"new_password" binding:"required,min=6,max=20"` // New Password of the User
}
