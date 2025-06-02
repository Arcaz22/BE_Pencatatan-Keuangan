package routes

import (
    "pencatatan_keuangan/internal/handler"
    "pencatatan_keuangan/internal/middleware"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/budget"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupBudgetRoutes(r *gin.Engine, db *gorm.DB) {
	budgetRepo := repository.NewBudgetRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	budgetSvc := budget.NewBudgetService(budgetRepo, categoryRepo)
	budgetHandler := handler.NewBudgetHandler(budgetSvc)

	budgetGroup := r.Group("/budgets")
	{
		authorized := budgetGroup.Group("")
		authorized.Use(middleware.AuthMiddleware())
		{
			authorized.POST("/create", budgetHandler.Create)
			authorized.GET("/all", budgetHandler.GetAll)
			authorized.PUT("/:id", budgetHandler.Update)
			authorized.DELETE("/:id", budgetHandler.Delete)
		}
	}
}
