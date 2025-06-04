package routes

import (
    "pencatatan_keuangan/internal/handler"
    "pencatatan_keuangan/internal/middleware"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/dashboard"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupDashboardRoutes(r *gin.Engine, db *gorm.DB) {
    dashboardRepo := repository.NewDashboardRepository(db)
    dashboardSvc := dashboard.NewDashboardService(dashboardRepo)
    dashboardHandler := handler.NewDashboardHandler(dashboardSvc)

    dashboardGroup := r.Group("/dashboard")
    {
        dashboardGroup.Use(middleware.AuthMiddleware())
        dashboardGroup.GET("", dashboardHandler.GetDashboard)
    }
}
