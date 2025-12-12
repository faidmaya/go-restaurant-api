package controllers

import (
	"net/http"
	"restaurant-api/models"
	"restaurant-api/repositories"
	"restaurant-api/tools"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserRepo *repositories.UserRepo
}

func NewAuthController(repo *repositories.UserRepo) *AuthController {
	return &AuthController{UserRepo: repo}
}

// REGISTER
func (c *AuthController) Register(ctx *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Role == "" {
		req.Role = "customer"
	}

	// === CEK EMAIL SUDAH ADA ATAU BELUM ===
	existing, _ := c.UserRepo.GetByEmail(req.Email)
	if existing != nil && existing.ID != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "email sudah terdaftar",
		})
		return
	}

	// Hash password
	hashed, _ := tools.HashPassword(req.Password)

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
		Role:     req.Role,
	}

	created, err := c.UserRepo.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered",
		"user":    created,
	})
}

// LOGIN
func (c *AuthController) Login(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek user ada atau tidak
	user, err := c.UserRepo.GetByEmail(req.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// cek password
	if !tools.CheckPassword(req.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// generate JWT
	token, err := tools.GenerateJWT(user.ID, user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"token":   token,
		"user": gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"role":       user.Role,
			"created_at": user.CreatedAt,
		},
	})
}
