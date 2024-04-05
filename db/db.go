package db

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
	SSLMode  string
}

func InitDB(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s", config.Username, config.Password, config.DBName, config.Host, config.Port, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("error Connecting to database: %s", err)
	}

	return db
}
