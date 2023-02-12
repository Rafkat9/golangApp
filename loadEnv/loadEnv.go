package loadEnv

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	///////////////////////////////////////.env connect
	log.Println(".env connect")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
