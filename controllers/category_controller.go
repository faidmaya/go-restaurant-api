package controllers

import (
	"net/http"

	"restaurant-api/models"
	"restaurant-api/repositories"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	Repo *repositories.CategoryRepo
}

func NewCategoryController(r *repositories.CategoryRepo) *CategoryController {
	return &CategoryController{Repo: r}
}

func (cc *CategoryController) Create(c *gin.Context) {
	var in models.Category
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.Repo.Create(&in); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, in)
}

func (cc *CategoryController) List(c *gin.Context) {
	out, err := cc.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}

func (cc *CategoryController) Update(c *gin.Context) {
	id := c.Param("id")
	var in struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.Repo.Update(id, in.Name, in.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category updated"})
}

func (cc *CategoryController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := cc.Repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
