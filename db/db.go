package db

import (
	"fmt"

	"github.com/2k4sm/httpCoffee/models"
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

func InitDB(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s", config.Username, config.Password, config.DBName, config.Host, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error Connecting to database: %s", err)
	}
	log.Infof("DB:%s successfully connected", config.DBName)

	autoMigrateDB(db)

	return db
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

	log.Infof("automigration of models completed.")
}
