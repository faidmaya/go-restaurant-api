package controllers

import (
	"net/http"

	"restaurant-api/models"
	"restaurant-api/repositories"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	Repo *repositories.MenuRepo
}

func NewMenuController(r *repositories.MenuRepo) *MenuController { return &MenuController{Repo: r} }

func (mc *MenuController) Create(c *gin.Context) {
	var in models.Menu
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := mc.Repo.Create(&in); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, in)
}

func (mc *MenuController) List(c *gin.Context) {
	out, err := mc.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}
