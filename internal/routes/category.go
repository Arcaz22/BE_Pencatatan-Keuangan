package routes

import (
	"pencatatan_keuangan/internal/handler"
	"pencatatan_keuangan/internal/middleware"
	"pencatatan_keuangan/internal/repository"
	"pencatatan_keuangan/internal/service/category"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCategoryRoutes(r *gin.Engine, db *gorm.DB) {
	categoryRepo := repository.NewCategoryRepository(db)
	categorySvc := category.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categorySvc)

	categoryGroup := r.Group("/categories")
	{
		authorized := categoryGroup.Group("")
		authorized.Use(middleware.AuthMiddleware())
		{
			authorized.POST("/create", categoryHandler.Create)
			authorized.GET("/all", categoryHandler.GetAll)
			authorized.PUT("/:id", categoryHandler.Update)
    		authorized.DELETE("/:id", categoryHandler.Delete)
		}
	}
}
