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
// @Summary Register a new user
// @Description Register a new user with username, password, and email
// @Tags user
// @Accept json
// @Produce json
// @Param user body dto.RegisterUserReq true "User registration details"
// @Success 201 {object} dto.APIResponse "User registered successfully"
// @Failure 400 {object} dto.APIResponse "Invalid request parameters"
// @Failure 500 {object} dto.APIResponse "Failed to register user"
// @Router /user/register [post]
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
// @Summary Login a user
// @Description Login a user with username and Password
// @Tags user
// @Accept json
// @Produce json
// @Param user body dto.LoginUserReq true "User login details"
// @Success 200 {object} dto.APIResponse "User logged in successfully"
// @Failure 400 {object} dto.APIResponse "Invalid request parameters"
// @Failure 500 {object} dto.APIResponse "Failed to login user"
// @Router /user/login [post]
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
// @Summary Logout a user
// @Description Logout a user by clearing their session
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} dto.APIResponse "User logged out successfully"
// @Failure 400 {object} dto.APIResponse "User not found"
// @Failure 500 {object} dto.APIResponse "Failed to logout user"
// @Router /user/logout [post]
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
// @Summary Reset user Password
// @Description Reset user password with username and new Password
// @Tags user
// @Accept json
// @Produce json
// @Param user body dto.ResetPasswordReq true "User password reset details"
// @Success 200 {object} dto.APIResponse "Password reset successfully"
// @Failure 400 {object} dto.APIResponse "Invalid request parameters"
// @Failure 500 {object} dto.APIResponse "Failed to reset password"
// @Router /user/reset-password [post]
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
