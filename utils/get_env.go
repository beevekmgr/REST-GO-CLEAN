package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Get value from env
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

//Load .env values
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error occured .env")
	}
}
