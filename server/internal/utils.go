package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func FetchEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading Environment variables: ", err)
	}
}
