package config

import (
	"fmt"
	"log"
	"pencatatan_keuangan/internal/domain"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
    DBHost     string `mapstructure:"DB_HOST"`
    DBPort     string `mapstructure:"DB_PORT"`
    DBUser     string `mapstructure:"DB_USER"`
    DBPassword string `mapstructure:"DB_PASSWORD"`
    DBName     string `mapstructure:"DB_NAME"`
    ServerPort string `mapstructure:"SERVER_PORT"`
}

func LoadConfig() (Config, error) {
    var config Config

    v := viper.New()

    v.AutomaticEnv()

    v.SetDefault("SERVER_PORT", "8080")
    v.SetDefault("DB_HOST", "localhost")
    v.SetDefault("DB_PORT", "5432")
    v.SetDefault("DB_USER", "postgres")
    v.SetDefault("DB_NAME", "pencatatan_keuangan")

    v.SetConfigFile(".env")
    if err := v.ReadInConfig(); err != nil {
        log.Println("No .env file found or error reading it. Using environment variables and defaults.")
    }

    if err := v.Unmarshal(&config); err != nil {
        return Config{}, fmt.Errorf("failed to unmarshal configuration: %w", err)
    }

    return config, nil
}

func InitDB(config Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    log.Printf("Connected to database at %s:%s", config.DBHost, config.DBPort)

    // Run migrations before returning the database connection
    if err := RunMigrations(db); err != nil {
        return nil, fmt.Errorf("failed to run database migrations: %w", err)
    }

    return db, nil
}

func RunMigrations(db *gorm.DB) error {
    log.Println("Running database migrations...")
    startTime := time.Now()

    err := db.AutoMigrate(
        &domain.User{},
        // &domain.Transaction{},
        // &domain.Category{},
    )

    if err != nil {
        log.Printf("Migration failed: %v", err)
        return err
    }

    log.Printf("Migration completed successfully in %v", time.Since(startTime))
    return nil
}
