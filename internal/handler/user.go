package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/n3xtchen/gin-3at/internal/dto"
	"github.com/n3xtchen/gin-3at/internal/service"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		Service: userService,
	}
}

// RegisterUser handles user registration requests.
func (user *UserHandler) RegisterUser(c *gin.Context) {
	var userReq dto.RegisterUserReq
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request parameters"})
		return
	}

	err := user.Service.Register(userReq.ToEntity())
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(201, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login requests.
func (user *UserHandler) LoginUser(c *gin.Context) {
	var loginReq dto.LoginUserReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request parameters"})
		return
	}

	userEntity, err := user.Service.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to login user"})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "user": dto.ToLoginUserRes(userEntity)})
}

// LogoutUser handles user logout requests.
func (user *UserHandler) LogoutUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	err := user.Service.Logout(userID.(int))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to logout user"})
		return
	}

	c.JSON(200, gin.H{"message": "User logged out successfully"})
}

// ResetPassword handles user password reset requests.
func (user *UserHandler) ResetPassword(c *gin.Context) {
	var resetReq dto.ResetPasswordReq
	if err := c.ShouldBindJSON(&resetReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request parameters"})
		return
	}

	err := user.Service.ResetPassword(resetReq.Username, resetReq.NewPassword)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to reset password"})
		return
	}

	c.JSON(200, gin.H{"message": "Password reset successfully"})
}
