package routers

import (
	"database/sql"
	"net/http"

	"restaurant-api/controllers"
	"restaurant-api/middlewares"
	"restaurant-api/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// ====== REPOSITORIES ======
	userRepo := repositories.NewUserRepo(db)
	categoryRepo := repositories.NewCategoryRepo(db)
	menuRepo := repositories.NewMenuRepo(db)
	orderRepo := repositories.NewOrderRepo(db)

	// ====== CONTROLLERS ======
	authC := controllers.NewAuthController(userRepo)
	categoryC := controllers.NewCategoryController(categoryRepo)
	menuC := controllers.NewMenuController(menuRepo)
	orderC := controllers.NewOrderController(orderRepo)

	// ====== PUBLIC ROUTES ======
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API running"})
	})

	api := r.Group("/api")
	{
		api.POST("/users/register", authC.Register)
		api.POST("/users/login", authC.Login)

		api.GET("/categories", categoryC.List)
		api.GET("/menus", menuC.List)
	}

	// ====== ADMIN ROUTES (JWT) ======
	admin := r.Group("/admin")
	admin.Use(middlewares.JWTMiddleware())
	{
		admin.POST("/categories", categoryC.Create)
		admin.POST("/menus", menuC.Create)
		admin.PUT("/categories/:id", categoryC.Update)
		admin.PUT("/menus/:id", menuC.Update)
		admin.DELETE("/categories/:id", categoryC.Delete)
		admin.DELETE("/menus/:id", menuC.Delete)
	}

	// ====== SECURE ROUTES (JWT) ======
	secure := r.Group("/secure")
	secure.Use(middlewares.JWTMiddleware())
	{
		secure.POST("/orders", orderC.Create)
		secure.GET("/orders", orderC.ListByUser)
	}

	return r
}
