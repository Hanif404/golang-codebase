package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err := godotenv.Load(filepath.Join(pwd, ".env")); err != nil {
		log.Println("⚠️  .env file not found. Using system env variables")
	}
}

func GetEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
