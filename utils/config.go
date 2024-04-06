package utils

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func GetEnv(field string) string {
	env := os.Getenv(field)

	if env == "" {
		log.Infof("no env variables with name: %s found", field)
	}

	return env
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Warnf("env variables not loaded from .env file : %s", err)
	}
}
