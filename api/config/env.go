package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadFromEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading the env file")
	}

	return os.Getenv(key)
}
