package config

import (
	"log"

	"github.com/joho/godotenv"
)

// load .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
