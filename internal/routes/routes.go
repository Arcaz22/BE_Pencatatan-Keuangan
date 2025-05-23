package routes

import (
    "net/http"
    "pencatatan_keuangan/internal/middleware"
    "pencatatan_keuangan/internal/routes/user"

    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    r.Use(middleware.CORS())
    r.Use(middleware.Logger())
    r.Use(middleware.ErrorHandler())

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.GET("/health", healthCheck)

    user.SetupRoutes(r, db)

    return r
}

// @Summary Health check endpoint
// @Description Check if the API is running
// @Tags health
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func healthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "ok",
        "message": "API is running",
    })
}
