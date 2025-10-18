package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort     string
	DatabaseURL string
}

func Load() *Config {
	// err := godotenv.Load(".env", "../.env", "../../.env")
	// if err != nil {
	// 	log.Println("Warning: .env file not found, using environment variables.")
	// }

	loadEnv()

	return &Config{
		AppPort:     os.Getenv("APP_PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + "user" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
