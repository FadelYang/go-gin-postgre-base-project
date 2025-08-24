package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using default environtment variables")
	}
}
