package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func EnvLineAccessToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading Access Token")
	}

	return os.Getenv("LINEACCESSTOKEN")
}