package routes

import (
    "pencatatan_keuangan/internal/handler"
    "pencatatan_keuangan/internal/middleware"
    "pencatatan_keuangan/internal/repository"
    "pencatatan_keuangan/internal/service/user"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
    userRepo := repository.NewUserRepository(db)
    userSvc := user.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userSvc)

    userGroup := r.Group("/users")
    {
        userGroup.POST("/register", userHandler.Register)
        userGroup.POST("/signin", userHandler.SignIn)

        authorized := userGroup.Group("")
        authorized.Use(middleware.AuthMiddleware())
        {
            authorized.GET("/profile", userHandler.Profile)
        }
    }
}
