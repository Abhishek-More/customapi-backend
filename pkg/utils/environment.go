package utils

/*
Loads environment variables as declared in the .env file
*/

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

}