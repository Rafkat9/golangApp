package postgres

import (
	"database/sql"
	"log"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	log.Println("db connect")
	var err error

	log.Println("create const env database")

	dsn := os.Getenv("DB_URL_POSTGRES")
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	defer DB.Close()

}
