package main

import (
    "log"
    "pencatatan_keuangan/config"
    "pencatatan_keuangan/docs"
    "pencatatan_keuangan/internal/routes"
)

// @title           Pencatatan Keuangan API
// @version         1.0
// @description     API for financial record keeping
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
    docs.SwaggerInfo.Title = "Pencatatan Keuangan API"
    docs.SwaggerInfo.Description = "API for financial recording application"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:8080"
    docs.SwaggerInfo.BasePath = "/"
    docs.SwaggerInfo.Schemes = []string{"http", "https"}

    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    db, err := config.InitDB(cfg)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    r := routes.SetupRouter(db)

    log.Printf("Server starting on port %s", cfg.ServerPort)
    if err := r.Run(":" + cfg.ServerPort); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
