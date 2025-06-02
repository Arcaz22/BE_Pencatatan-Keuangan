package routes

import (
    "pencatatan_keuangan/internal/handler"
    "pencatatan_keuangan/internal/middleware"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/expense"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupExpenseRoutes(r *gin.Engine, db *gorm.DB) {
    expenseRepo := repository.NewExpenseRepository(db)
    categoryRepo := repository.NewCategoryRepository(db)
    expenseSvc := expense.NewExpenseService(expenseRepo, categoryRepo)
    expenseHandler := handler.NewExpenseHandler(expenseSvc)

    expenseGroup := r.Group("/expenses")
    {
        authorized := expenseGroup.Group("")
        authorized.Use(middleware.AuthMiddleware())
        {
            authorized.POST("/create", expenseHandler.Create)
			authorized.GET("/all", expenseHandler.GetAll)
			authorized.PUT("/:id", expenseHandler.Update)
        	authorized.DELETE("/:id", expenseHandler.Delete)
        }
    }
}
