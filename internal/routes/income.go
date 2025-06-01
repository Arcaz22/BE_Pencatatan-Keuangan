package routes

import (
    "pencatatan_keuangan/internal/handler"
    "pencatatan_keuangan/internal/middleware"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/income"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupIncomeRoutes(r *gin.Engine, db *gorm.DB) {
    incomeRepo := repository.NewIncomeRepository(db)
    categoryRepo := repository.NewCategoryRepository(db)
    incomeSvc := income.NewIncomeService(incomeRepo, categoryRepo)
    incomeHandler := handler.NewIncomeHandler(incomeSvc)

    incomeGroup := r.Group("/incomes")
    {
        authorized := incomeGroup.Group("")
        authorized.Use(middleware.AuthMiddleware())
        {
            authorized.POST("/create", incomeHandler.Create)
			authorized.GET("/all", incomeHandler.GetAll)
			authorized.PUT("/:id", incomeHandler.Update)
        	authorized.DELETE("/:id", incomeHandler.Delete)
        }
    }
}
