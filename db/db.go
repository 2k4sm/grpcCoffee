package db

import (
	"fmt"

	models "github.com/2k4sm/httpCoffee/entities"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

var DB *gorm.DB

func InitDB(config *Config) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s", config.Username, config.Password, config.DBName, config.Host, config.Port)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error Connecting to database: %s", err)
	}

	autoMigrateDB(DB)
}

func autoMigrateDB(db *gorm.DB) {
	userModel := models.User{}
	coffeeModel := models.Coffee{}
	paymentModel := models.Payment{}
	coffeeHouseModel := models.CoffeeHouse{}

	err := db.AutoMigrate(userModel, coffeeModel, paymentModel, coffeeHouseModel)
	if err != nil {
		log.Warnf("error automigrating models :%s", err)
	}

}
